/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "..\..\fetch.pb"
import * as GoogleProtobufEmpty from "..\..\google\protobuf\empty.pb"
import * as WhoisV1Whois from "..\..\whois\v1\whois.pb"
export class NodeManagerService {
  static GetBotsInfo(req: GoogleProtobufEmpty.Empty, entityNotifier?: fm.NotifyStreamEntityArrival<WhoisV1Whois.WhoisInfo>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<GoogleProtobufEmpty.Empty, WhoisV1Whois.WhoisInfo>(`/v1/node/bots?${fm.renderURLSearchParams(req, [])}`, entityNotifier, {...initReq, method: "GET"})
  }
}