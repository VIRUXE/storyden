/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Arguments, Key, SWRConfiguration } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";

import type {
  CollectionAddPostOKResponse,
  CollectionCreateBody,
  CollectionCreateOKResponse,
  CollectionGetOKResponse,
  CollectionListOKResponse,
  CollectionRemovePostOKResponse,
  CollectionUpdateBody,
  CollectionUpdateOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Create a collection for curating posts under the authenticated account.

 */
export const collectionCreate = (
  collectionCreateBody: CollectionCreateBody,
) => {
  return fetcher<CollectionCreateOKResponse>({
    url: `/v1/collections`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: collectionCreateBody,
  });
};

export const getCollectionCreateMutationFetcher = () => {
  return (
    _: string,
    { arg }: { arg: CollectionCreateBody },
  ): Promise<CollectionCreateOKResponse> => {
    return collectionCreate(arg);
  };
};
export const getCollectionCreateMutationKey = () => `/v1/collections` as const;

export type CollectionCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionCreate>>
>;
export type CollectionCreateMutationError =
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useCollectionCreate = <
  TError = UnauthorisedResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof collectionCreate>>,
    TError,
    string,
    CollectionCreateBody,
    Awaited<ReturnType<typeof collectionCreate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getCollectionCreateMutationKey();
  const swrFn = getCollectionCreateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * List all collections using the filtering options.
 */
export const collectionList = () => {
  return fetcher<CollectionListOKResponse>({
    url: `/v1/collections`,
    method: "GET",
  });
};

export const getCollectionListKey = () => [`/v1/collections`] as const;

export type CollectionListQueryResult = NonNullable<
  Awaited<ReturnType<typeof collectionList>>
>;
export type CollectionListQueryError =
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof collectionList>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getCollectionListKey() : null));
  const swrFn = () => collectionList();

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
 * Get a collection by its ID. Collections can be public or private so the
response will depend on which account is making the request and if the
target collection is public, private, owned or not owned by the account.

 */
export const collectionGet = (collectionId: string) => {
  return fetcher<CollectionGetOKResponse>({
    url: `/v1/collections/${collectionId}`,
    method: "GET",
  });
};

export const getCollectionGetKey = (collectionId: string) =>
  [`/v1/collections/${collectionId}`] as const;

export type CollectionGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof collectionGet>>
>;
export type CollectionGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof collectionGet>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!collectionId;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getCollectionGetKey(collectionId) : null));
  const swrFn = () => collectionGet(collectionId);

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
 * Update a collection owned by the authenticated account.
 */
export const collectionUpdate = (
  collectionId: string,
  collectionUpdateBody: CollectionUpdateBody,
) => {
  return fetcher<CollectionUpdateOKResponse>({
    url: `/v1/collections/${collectionId}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: collectionUpdateBody,
  });
};

export const getCollectionUpdateMutationFetcher = (collectionId: string) => {
  return (
    _: string,
    { arg }: { arg: CollectionUpdateBody },
  ): Promise<CollectionUpdateOKResponse> => {
    return collectionUpdate(collectionId, arg);
  };
};
export const getCollectionUpdateMutationKey = (collectionId: string) =>
  `/v1/collections/${collectionId}` as const;

export type CollectionUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionUpdate>>
>;
export type CollectionUpdateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionUpdate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionUpdate>>,
      TError,
      string,
      CollectionUpdateBody,
      Awaited<ReturnType<typeof collectionUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getCollectionUpdateMutationKey(collectionId);
  const swrFn = getCollectionUpdateMutationFetcher(collectionId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add a post to a collection. The collection must be owned by the account
making the request. The post can be any published post of any kind.

 */
export const collectionAddPost = (collectionId: string, postId: string) => {
  return fetcher<CollectionAddPostOKResponse>({
    url: `/v1/collections/${collectionId}/posts/${postId}`,
    method: "PUT",
  });
};

export const getCollectionAddPostMutationFetcher = (
  collectionId: string,
  postId: string,
) => {
  return (
    _: string,
    __: { arg: Arguments },
  ): Promise<CollectionAddPostOKResponse> => {
    return collectionAddPost(collectionId, postId);
  };
};
export const getCollectionAddPostMutationKey = (
  collectionId: string,
  postId: string,
) => `/v1/collections/${collectionId}/posts/${postId}` as const;

export type CollectionAddPostMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionAddPost>>
>;
export type CollectionAddPostMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionAddPost = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionId: string,
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionAddPost>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof collectionAddPost>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getCollectionAddPostMutationKey(collectionId, postId);
  const swrFn = getCollectionAddPostMutationFetcher(collectionId, postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove a post from a collection. The collection must be owned by the
account making the request.

 */
export const collectionRemovePost = (collectionId: string, postId: string) => {
  return fetcher<CollectionRemovePostOKResponse>({
    url: `/v1/collections/${collectionId}/posts/${postId}`,
    method: "DELETE",
  });
};

export const getCollectionRemovePostMutationFetcher = (
  collectionId: string,
  postId: string,
) => {
  return (
    _: string,
    __: { arg: Arguments },
  ): Promise<CollectionRemovePostOKResponse> => {
    return collectionRemovePost(collectionId, postId);
  };
};
export const getCollectionRemovePostMutationKey = (
  collectionId: string,
  postId: string,
) => `/v1/collections/${collectionId}/posts/${postId}` as const;

export type CollectionRemovePostMutationResult = NonNullable<
  Awaited<ReturnType<typeof collectionRemovePost>>
>;
export type CollectionRemovePostMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useCollectionRemovePost = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  collectionId: string,
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof collectionRemovePost>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof collectionRemovePost>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getCollectionRemovePostMutationKey(collectionId, postId);
  const swrFn = getCollectionRemovePostMutationFetcher(collectionId, postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
