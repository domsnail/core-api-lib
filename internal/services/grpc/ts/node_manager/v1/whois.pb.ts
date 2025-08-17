/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "..\..\fetch.pb"
import * as WhoisV1Whois from "..\..\whois\v1\whois.pb"
export class WhoisService {
  static NewProbe(req: WhoisV1Whois.ProbeOptions, entityNotifier?: fm.NotifyStreamEntityArrival<WhoisV1Whois.WhoisInfo>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<WhoisV1Whois.ProbeOptions, WhoisV1Whois.WhoisInfo>(`/v1/whois/probe`, entityNotifier, {...initReq, method: "POST"})
  }
}