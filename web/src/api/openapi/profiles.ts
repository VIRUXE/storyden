/**
 * Generated by orval v6.24.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";

import type {
  InternalServerErrorResponse,
  NotFoundResponse,
  ProfileGetOKResponse,
  ProfileListOKResponse,
  ProfileListParams,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Query and search profiles.
 */
export const profileList = (params?: ProfileListParams) => {
  return fetcher<ProfileListOKResponse>({
    url: `/v1/profiles`,
    method: "GET",
    params,
  });
};

export const getProfileListKey = (params?: ProfileListParams) =>
  [`/v1/profiles`, ...(params ? [params] : [])] as const;

export type ProfileListQueryResult = NonNullable<
  Awaited<ReturnType<typeof profileList>>
>;
export type ProfileListQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useProfileList = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params?: ProfileListParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof profileList>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getProfileListKey(params) : null));
  const swrFn = () => profileList(params);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Get a public profile by ID.
 */
export const profileGet = (accountHandle: string) => {
  return fetcher<ProfileGetOKResponse>({
    url: `/v1/profiles/${accountHandle}`,
    method: "GET",
  });
};

export const getProfileGetKey = (accountHandle: string) =>
  [`/v1/profiles/${accountHandle}`] as const;

export type ProfileGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof profileGet>>
>;
export type ProfileGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useProfileGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  accountHandle: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof profileGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!accountHandle;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getProfileGetKey(accountHandle) : null));
  const swrFn = () => profileGet(accountHandle);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
