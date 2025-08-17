/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "..\..\google\protobuf\timestamp.pb"
import * as HostsV1Host from "..\..\hosts\v1\host.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type ProbeOptions = {
  responseRequired?: boolean
  hosts?: HostsV1Host.Host[]
}


type BaseWhoisInfo = {
  host?: string
  rawData?: string
  sourceServer?: string
  createdAt?: GoogleProtobufTimestamp.Timestamp
}

export type WhoisInfo = BaseWhoisInfo
  & OneOf<{ redirected: WhoisInfo }>