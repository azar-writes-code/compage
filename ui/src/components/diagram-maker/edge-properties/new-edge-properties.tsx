/* tslint:disable:no-string-literal */
import React, {ChangeEvent} from "react";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import TextField from "@mui/material/TextField";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import {getCurrentConfig, setModifiedState} from "../../../utils/localstorage-client";
import {getParsedModifiedState} from "../helper/helper";
import Divider from "@mui/material/Divider";
import {Checkbox, FormControlLabel, Stack} from "@mui/material";
import {CompageEdge, CompageNode, GrpcClient, GrpcConfig, RestClient, RestConfig, WsClient, WsConfig} from "../models";
import {updateModifiedState} from "../../../features/projects-operations/populateModifiedState";
import {COMPAGE} from "../node-properties/utils";

interface NewEdgePropertiesProps {
    isOpen: boolean;
    edgeId: string;
    onEdgePropertiesClose: () => void;
}

interface ClientTypesConfig {
    isRestServer?: boolean;
    restServerPort?: string;
    isGrpcServer?: boolean;
    grpcServerPort?: string;
    isWsServer?: boolean;
    wsServerPort?: string;
    sourceNodeName?: string;
}

const getClientTypesConfig = (parsedCurrentConfig, parsedModifiedState, edgeId): ClientTypesConfig => {
    // read node's clients array and get these values
    const clientTypesConfig: ClientTypesConfig = {};
    const edgeConfig: CompageEdge = parsedCurrentConfig.edges[edgeId];
    const destNode: CompageNode = parsedModifiedState.nodes[edgeConfig?.dest];
    const srcNode: CompageNode = parsedModifiedState.nodes[edgeConfig?.src];
    if (srcNode && destNode) {
        // rest - extract clients in dest node
        const restClients: RestClient[] = destNode?.consumerData?.restConfig?.clients;
        if (restClients && restClients.length > 0) {
            for (const restClient of restClients) {
                if (srcNode.consumerData.name === restClient.sourceNodeName) {
                    clientTypesConfig.isRestServer = true;
                    clientTypesConfig.restServerPort = srcNode?.consumerData?.restConfig?.server?.port;
                }
            }
        }
        // grpc - extract clients in dest node
        const grpcClients: GrpcClient[] = destNode?.consumerData?.grpcConfig?.clients;
        if (grpcClients && grpcClients.length > 0) {
            for (const grpcClient of grpcClients) {
                if (srcNode.consumerData.name === grpcClient.sourceNodeName) {
                    clientTypesConfig.isGrpcServer = true;
                    clientTypesConfig.grpcServerPort = srcNode?.consumerData?.grpcConfig?.server?.port;
                }
            }
        }
        // ws - extract clients in dest node
        const wsClients: WsClient[] = destNode?.consumerData?.wsConfig?.clients;
        if (wsClients && wsClients.length > 0) {
            for (const wsClient of wsClients) {
                if (srcNode.consumerData.name === wsClient.sourceNodeName) {
                    clientTypesConfig.isWsServer = true;
                    clientTypesConfig.wsServerPort = srcNode?.consumerData?.wsConfig?.server?.port;
                }
            }
        }
    }
    return clientTypesConfig;
};

export const NewEdgeProperties = (props: NewEdgePropertiesProps) => {
    const parsedCurrentConfig = JSON.parse(getCurrentConfig());
    let parsedModifiedState = getParsedModifiedState();
    // sometimes the parsedModifiedState is empty so, recreate it.
    if (Object.keys(parsedModifiedState.edges).length < 1) {
        updateModifiedState(JSON.parse(getCurrentConfig()));
        parsedModifiedState = getParsedModifiedState();
    }
    const modifiedEdgeState: CompageEdge = parsedModifiedState.edges[props.edgeId];

    const clientTypesConfig: ClientTypesConfig = getClientTypesConfig(parsedCurrentConfig, parsedModifiedState, props.edgeId);

    const [payload, setPayload] = React.useState({
        name: modifiedEdgeState?.consumerData?.name !== undefined ? modifiedEdgeState?.consumerData?.name : "",
        isRestServer: clientTypesConfig.isRestServer || false,
        restServerPort: clientTypesConfig.restServerPort,
        isGrpcServer: clientTypesConfig.isGrpcServer || false,
        grpcServerPort: clientTypesConfig.grpcServerPort,
        isWsServer: clientTypesConfig.isWsServer || false,
        wsServerPort: clientTypesConfig.wsServerPort,
        sourceNodeName: clientTypesConfig.sourceNodeName,
    });

    // TODO this is a hack as there is no EDGE_UPDATE action in diagram-maker. We may later update this impl when we fork diagram-maker repo.
    // update state with additional properties added from UI (Post edge creation)
    const handleEdgeUpdate = (event: React.MouseEvent<HTMLElement>) => {
        event.preventDefault();
        const modifiedState = getParsedModifiedState();
        // update modifiedState with current fields on dialog box
        // P.S. - We will have the fields in consumerData which are on dialogBox, so we can assign them directly. We also refer the older values when payload state is initialized, so the older values will be persisted as they are if not changed.
        if (!(props.edgeId in modifiedState.edges)) {
            // adding consumerData to new edge in modifiedState
            modifiedState.edges[props.edgeId] = {
                consumerData: {},
            };
        }

        // referring edge from config here instead of modifiedState as the modifiedState doesn't have a src and dest.
        const edgeConfig: CompageEdge = parsedCurrentConfig.edges[props.edgeId];
        const modifiedEdgeState: CompageEdge = modifiedState.edges[props.edgeId];
        const dstNode: CompageNode = modifiedState.nodes[edgeConfig.dest];
        modifiedEdgeState.consumerData.name = payload.name;
        // get dest node and add details to it.
        if (payload.isRestServer) {
            const restClient: RestClient = {
                sourceNodeName: payload.sourceNodeName,
                port: payload.restServerPort,
            };
            for (let i = 0; i < dstNode?.consumerData?.restConfig?.clients.length; i++) {
                // search for old restClient and delete it.
                if (dstNode.consumerData?.restConfig?.clients[i].sourceNodeName === payload.sourceNodeName) {
                    dstNode.consumerData?.restConfig?.clients.splice(i, 1);
                    break;
                }
            }
            dstNode?.consumerData?.restConfig?.clients.push(restClient);
        }
        if (payload.isGrpcServer) {
            const grpcClient: GrpcClient = {
                sourceNodeName: payload.sourceNodeName,
                port: payload.grpcServerPort,
            };
            debugger
            for (let i = 0; i < dstNode?.consumerData?.grpcConfig?.clients?.length; i++) {
                // search for old grpcClient and delete it.
                if (dstNode?.consumerData?.grpcConfig?.clients[i]?.sourceNodeName === payload.sourceNodeName) {
                    dstNode?.consumerData?.grpcConfig?.clients.splice(i, 1);
                    break;
                }
            }
            if (dstNode?.consumerData?.grpcConfig) {
                dstNode?.consumerData?.grpcConfig?.clients.push(grpcClient);
            } else {
                // The port had to be made optional in server here. The below code will be in picture when you don't have
                // grpc configured in dst node
                dstNode.consumerData.grpcConfig = {template: COMPAGE, clients: [], server: {}};
                // push the client now.
                dstNode?.consumerData?.grpcConfig?.clients.push(grpcClient);
            }

        }
        if (payload.isWsServer) {
            const wsClient: WsClient = {
                sourceNodeName: payload.sourceNodeName,
                port: payload.wsServerPort,
            };
            for (let i = 0; i < dstNode?.consumerData?.wsConfig?.clients.length; i++) {
                // search for old wsClient and delete it.
                if (dstNode?.consumerData?.wsConfig?.clients[i].sourceNodeName === payload.sourceNodeName) {
                    dstNode?.consumerData?.wsConfig?.clients.splice(i, 1);
                    break;
                }
            }
            dstNode?.consumerData?.wsConfig?.clients.push(wsClient);
        }

        // update modifiedState in the localstorage
        setModifiedState(JSON.stringify(modifiedState));
        setPayload({
            name: "",
            isRestServer: false,
            restServerPort: "",
            isGrpcServer: false,
            grpcServerPort: "",
            isWsServer: false,
            wsServerPort: "",
            sourceNodeName: ""
        });
        props.onEdgePropertiesClose();
    };

    const onEdgePropertiesClose = (e: any, reason: "backdropClick" | "escapeKeyDown") => {
        // this prevents dialog box from closing.
        if (reason === "backdropClick" || reason === "escapeKeyDown") {
            return;
        }
        props.onEdgePropertiesClose();
    };

    const handleNameChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            name: event.target.value
        });
    };

    const handleIsRestServerChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPayload({
            ...payload,
            isRestServer: event.target.checked
        });
    };

    const getRestServerCheck = () => {
        return <React.Fragment>
            <FormControlLabel
                label="Rest Server"
                control={<Checkbox
                    size="medium" checked={payload.isRestServer}
                    onChange={handleIsRestServerChange}
                />}
            />
        </React.Fragment>;
    };

    const getRestServerPort = () => {
        if (payload.isRestServer) {
            // retrieve port from src node.
            const srcNodeId = parsedCurrentConfig.edges[props.edgeId].src;
            const srcNode: CompageNode = parsedModifiedState.nodes[srcNodeId];
            const restConfig: RestConfig = srcNode?.consumerData?.restConfig;
            if (restConfig && restConfig.server.port) {
                payload.restServerPort = restConfig.server.port;
                payload.sourceNodeName = srcNode.consumerData.name;
            } else {
                // this is a default port for every project generated by openapi-generator
                payload.restServerPort = "8080";
            }
            return <TextField
                required
                size="medium"
                margin="dense"
                id="restServerPort"
                label="Rest Server Port"
                type="text"
                disabled
                value={payload.restServerPort}
                variant="outlined"
            />;
        }
        return "";
    };

    const handleIsGrpcServerChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPayload({
            ...payload,
            isGrpcServer: event.target.checked
        });
    };

    const getGrpcServerCheck = () => {
        return <React.Fragment>
            <FormControlLabel
                label="Grpc Server"
                control={<Checkbox
                    size="medium" checked={payload.isGrpcServer}
                    onChange={handleIsGrpcServerChange}
                />}
            />
        </React.Fragment>;
    };

    const getGrpcServerPort = () => {
        if (payload.isGrpcServer) {
            // retrieve port from src node.
            const srcNodeId: string = parsedCurrentConfig.edges[props.edgeId].src;
            const srcNode: CompageNode = parsedModifiedState.nodes[srcNodeId];
            const grpcConfig: GrpcConfig = srcNode?.consumerData?.grpcConfig;
            if (grpcConfig && Object.keys(grpcConfig).length > 0) {
                payload.grpcServerPort = grpcConfig.server.port;
            }
            return <TextField
                required
                size="medium"
                margin="dense"
                id="grpcServerPort"
                label="Grpc Server Port"
                type="text"
                disabled
                value={payload.grpcServerPort}
                variant="outlined"
            />;
        }
        return "";
    };

    const handleIsWsServerChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPayload({
            ...payload,
            isWsServer: event.target.checked
        });
    };

    const getWsServerCheck = () => {
        return <React.Fragment>
            <FormControlLabel
                label="Ws Server"
                disabled
                control={<Checkbox
                    size="medium" checked={payload.isWsServer}
                    onChange={handleIsWsServerChange}
                />}
            />
        </React.Fragment>;
    };

    const getWsServerPort = () => {
        if (payload.isWsServer) {
            // retrieve port from src node.
            const srcNode = parsedCurrentConfig.edges[props.edgeId].src;
            const wsConfig: WsConfig = parsedModifiedState.nodes[srcNode]?.consumerData?.wsConfig;
            if (wsConfig && Object.keys(wsConfig).length > 0) {
                payload.wsServerPort = wsConfig.server.port;
            }
            return <TextField
                required
                size="medium"
                margin="dense"
                id="wsServerPort"
                label="Port"
                type="text"
                value={payload.wsServerPort}
                variant="outlined"
            />;
        }
        return "";
    };

    return <React.Fragment>
        <Dialog open={props.isOpen} onClose={onEdgePropertiesClose}>
            <DialogTitle>Edge properties : {props.edgeId}</DialogTitle>
            <Divider/>
            <DialogContent>
                <Stack direction="column" spacing={2}>
                    <TextField
                        required
                        size="medium"
                        margin="dense"
                        id="name"
                        label="Name of edge"
                        type="text"
                        value={payload.name}
                        onChange={handleNameChange}
                        variant="outlined"
                    />
                    {getRestServerCheck()}
                    {getRestServerPort()}
                    {getGrpcServerCheck()}
                    {getGrpcServerPort()}
                    {/*{getWsServerCheck()}*/}
                    {/*{getWsServerPort()}*/}
                </Stack>
            </DialogContent>
            <DialogActions>
                <Button variant="outlined" color="secondary" onClick={props.onEdgePropertiesClose}>Cancel</Button>
                <Button variant="contained"
                        onClick={handleEdgeUpdate}
                        disabled={payload.name === ""}>Update</Button>
            </DialogActions>
        </Dialog>
    </React.Fragment>;
};
