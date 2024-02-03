/**
 * Generated by orval v6.24.0 🍺
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
  InternalServerErrorResponse,
  NotFoundResponse,
  PostCreateBody,
  PostCreateOKResponse,
  PostReactAddBody,
  PostReactAddOKResponse,
  PostSearchOKResponse,
  PostSearchParams,
  PostUpdateBody,
  PostUpdateOKResponse,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Create a new post within a thread.
 */
export const postCreate = (
  threadMark: string,
  postCreateBody: PostCreateBody,
) => {
  return fetcher<PostCreateOKResponse>({
    url: `/v1/threads/${threadMark}/posts`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: postCreateBody,
  });
};

export const getPostCreateMutationFetcher = (threadMark: string) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<PostCreateOKResponse> => {
    return postCreate(threadMark, arg as PostCreateBody);
  };
};
export const getPostCreateMutationKey = (threadMark: string) =>
  `/v1/threads/${threadMark}/posts` as const;

export type PostCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof postCreate>>
>;
export type PostCreateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostCreate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  threadMark: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postCreate>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof postCreate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostCreateMutationKey(threadMark);
  const swrFn = getPostCreateMutationFetcher(threadMark);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Publish changes to a single post.
 */
export const postUpdate = (postId: string, postUpdateBody: PostUpdateBody) => {
  return fetcher<PostUpdateOKResponse>({
    url: `/v1/posts/${postId}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: postUpdateBody,
  });
};

export const getPostUpdateMutationFetcher = (postId: string) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<PostUpdateOKResponse> => {
    return postUpdate(postId, arg as PostUpdateBody);
  };
};
export const getPostUpdateMutationKey = (postId: string) =>
  `/v1/posts/${postId}` as const;

export type PostUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof postUpdate>>
>;
export type PostUpdateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostUpdate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postUpdate>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof postUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostUpdateMutationKey(postId);
  const swrFn = getPostUpdateMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Archive a post using soft-delete.
 */
export const postDelete = (postId: string) => {
  return fetcher<void>({ url: `/v1/posts/${postId}`, method: "DELETE" });
};

export const getPostDeleteMutationFetcher = (postId: string) => {
  return (_: string, { arg }: { arg: Arguments }): Promise<void> => {
    return postDelete(postId);
  };
};
export const getPostDeleteMutationKey = (postId: string) =>
  `/v1/posts/${postId}` as const;

export type PostDeleteMutationResult = NonNullable<
  Awaited<ReturnType<typeof postDelete>>
>;
export type PostDeleteMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostDelete = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postDelete>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof postDelete>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostDeleteMutationKey(postId);
  const swrFn = getPostDeleteMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Search through posts using various queries and filters.
 */
export const postSearch = (params?: PostSearchParams) => {
  return fetcher<PostSearchOKResponse>({
    url: `/v1/posts/search`,
    method: "GET",
    params,
  });
};

export const getPostSearchKey = (params?: PostSearchParams) =>
  [`/v1/posts/search`, ...(params ? [params] : [])] as const;

export type PostSearchQueryResult = NonNullable<
  Awaited<ReturnType<typeof postSearch>>
>;
export type PostSearchQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostSearch = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params?: PostSearchParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof postSearch>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getPostSearchKey(params) : null));
  const swrFn = () => postSearch(params);

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
 * Add a reaction to a post.
 */
export const postReactAdd = (
  postId: string,
  postReactAddBody: PostReactAddBody,
) => {
  return fetcher<PostReactAddOKResponse>({
    url: `/v1/posts/${postId}/reacts`,
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    data: postReactAddBody,
  });
};

export const getPostReactAddMutationFetcher = (postId: string) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<PostReactAddOKResponse> => {
    return postReactAdd(postId, arg as PostReactAddBody);
  };
};
export const getPostReactAddMutationKey = (postId: string) =>
  `/v1/posts/${postId}/reacts` as const;

export type PostReactAddMutationResult = NonNullable<
  Awaited<ReturnType<typeof postReactAdd>>
>;
export type PostReactAddMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostReactAdd = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postReactAdd>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof postReactAdd>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostReactAddMutationKey(postId);
  const swrFn = getPostReactAddMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
