/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

type BaseHost = {
}

export type Host = BaseHost
  & OneOf<{ dn: string; url: string; email: string; ip: IPAddress }>
  & OneOf<{ uuid: string }>


type BaseIPAddress = {
}

export type IPAddress = BaseIPAddress
  & OneOf<{ v4: number; v6: Uint8Array }>