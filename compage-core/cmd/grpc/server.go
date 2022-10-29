package grpc

import (
	_ "embed"
	project "github.com/kube-tarian/compage-core/gen/api/v1"
	"github.com/kube-tarian/compage-core/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

var chunkSize = 1024 * 3

type server struct {
	project.UnimplementedProjectServiceServer
}

func New() project.ProjectServiceServer {
	return server{}
}

// CreateProject implements api.v1.CreateProject
func (s server) CreateProject(projectRequest *project.ProjectRequest, server project.ProjectService_CreateProjectServer) error {
	//projectGrpc, err := grpc.GetProject(projectRequest)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(projectGrpc.CompageYaml)
	// createProject
	err := utils.CreateTarFile(projectRequest.ProjectName, utils.GetProjectDirectoryName(projectRequest.GetProjectName()))
	if err != nil {
		return err
	}
	return sendFile(projectRequest, server)
}

// UpdateProject implements api.v1.UpdateProject
func (s server) UpdateProject(projectRequest *project.ProjectRequest, server project.ProjectService_UpdateProjectServer) error {
	//projectGrpc, err := grpc.GetProject(projectRequest)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(projectGrpc.CompageYaml)
	// createProject
	err := utils.CreateTarFile(projectRequest.ProjectName, utils.GetProjectDirectoryName(projectRequest.GetProjectName()))
	if err != nil {
		return err
	}
	return sendFile(projectRequest, server)
}

func sendFile(projectRequest *project.ProjectRequest, server project.ProjectService_CreateProjectServer) error {
	f, ok := utils.GetFile(utils.GetProjectTarFilePath(projectRequest.ProjectName))
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