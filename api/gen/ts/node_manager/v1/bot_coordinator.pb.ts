/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "..\..\fetch.pb"
export type BotCredentials = {
  apiKey?: string
  name?: string
}

export type ConnectionInfo = {
  queuePool?: string
  queuePassword?: string
}

export class BotCoordinatorService {
  static Connect(req: BotCredentials, initReq?: fm.InitReq): Promise<ConnectionInfo> {
    return fm.fetchReq<BotCredentials, ConnectionInfo>(`/whois.v1.BotCoordinatorService/Connect`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}