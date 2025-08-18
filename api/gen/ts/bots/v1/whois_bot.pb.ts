/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "..\..\fetch.pb"
import * as WhoisV1Whois from "..\..\whois\v1\whois.pb"
export class WhoisBotService {
  static ProbeHost(req: WhoisV1Whois.ProbeOptions, entityNotifier?: fm.NotifyStreamEntityArrival<WhoisV1Whois.WhoisInfo>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<WhoisV1Whois.ProbeOptions, WhoisV1Whois.WhoisInfo>(`/whois.v1.WhoisBotService/ProbeHost`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}