package grpc

import (
	_ "embed"
	project "github.com/kube-tarian/compage/core/gen/api/v1"
	"github.com/kube-tarian/compage/core/internal/converter/grpc"
	"github.com/kube-tarian/compage/core/internal/generator"
	"github.com/kube-tarian/compage/core/internal/taroperations"
	"github.com/kube-tarian/compage/core/internal/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
)

var chunkSize = 1024 * 3

type server struct {
	project.UnimplementedProjectServiceServer
}

func New() project.ProjectServiceServer {
	return &server{}
}

// CreateProject implements api.v1.CreateProject
func (s *server) CreateProject(projectRequest *project.ProjectRequest, server project.ProjectService_CreateProjectServer) error {
	// converts to core project
	coreProject, err := grpc.GetProject(projectRequest)
	if err != nil {
		log.Debug(err)
		return status.Errorf(codes.InvalidArgument,
			"error while converting request to project ["+err.Error()+"]")
	}

	// triggers project generation
	if err := generator.Generator(coreProject); err != nil {
		log.Debug(err)
		return status.Errorf(codes.InvalidArgument,
			"error while generating the project ["+err.Error()+"]")
	}

	// CreateTarFile creates tar file for the project generated
	err = taroperations.CreateTarFile(coreProject.Name, utils.GetProjectDirectoryName(projectRequest.GetProjectName()))
	if err != nil {
		log.Debug(err)
		return status.Errorf(codes.Internal,
			"error while converting request to project ["+err.Error()+"]")
	}

	// delete tmp/project-name folder
	defer func(name string) {
		if err = os.RemoveAll(name); err != nil {
			log.Debug(err)
		}
	}(utils.GetProjectDirectoryName(projectRequest.GetProjectName()))

	// delete just file
	defer func(name string) {
		if err = os.Remove(name); err != nil {
			log.Debug(err)
		}
	}(taroperations.GetProjectTarFilePath(projectRequest.GetProjectName()))

	// stream file to grpc client
	return sendFile(projectRequest, server)
}

// UpdateProject implements api.v1.UpdateProject
func (s *server) UpdateProject(projectRequest *project.ProjectRequest, server project.ProjectService_UpdateProjectServer) error {
	//projectGrpc, err := grpc.GetProject(projectRequest)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(projectGrpc.CompageYaml)
	// createProject
	err := taroperations.CreateTarFile(projectRequest.ProjectName, utils.GetProjectDirectoryName(projectRequest.GetProjectName()))
	if err != nil {
		log.Debug(err)
		return status.Errorf(codes.InvalidArgument,
			"error while creating a tar file ["+err.Error()+"]")
	}
	return sendFile(projectRequest, server)
}

func sendFile(projectRequest *project.ProjectRequest, server project.ProjectService_CreateProjectServer) error {
	f, ok := taroperations.GetFile(taroperations.GetProjectTarFilePath(projectRequest.ProjectName))
	if !ok {
		return status.Error(codes.NotFound, "file is not found")
	}
	err := server.SendHeader(f.Metadata())
	if err != nil {
		return status.Error(codes.Internal, "error during sending header")
	}
	fileChunk := &project.ProjectResponse{FileChunk: make([]byte, chunkSize)}
	var n int
Loop:
	for {
		n, err = f.Read(fileChunk.FileChunk)
		switch err {
		case nil:
		case io.EOF:
			break Loop
		default:
			return status.Errorf(codes.Internal, "io.ReadAll: %v", err)
		}
		fileChunk.FileChunk = fileChunk.FileChunk[:n]
		serverErr := server.Send(fileChunk)
		if serverErr != nil {
			return status.Errorf(codes.Internal, "server.Send: %v", serverErr)
		}
	}
	return nil
}