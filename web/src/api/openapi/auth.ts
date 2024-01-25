/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";

import type {
  AuthPasswordBody,
  AuthPasswordCreateBody,
  AuthPasswordUpdateBody,
  AuthProviderListOKResponse,
  AuthSuccessOKResponse,
  BadRequestResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  OAuthProviderCallbackBody,
  PhoneRequestCodeBody,
  PhoneSubmitCodeBody,
  UnauthorisedResponse,
  WebAuthnGetAssertionOKResponse,
  WebAuthnMakeAssertionBody,
  WebAuthnMakeCredentialBody,
  WebAuthnRequestCredentialOKResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Retrieve a list of authentication providers. Storyden supports a few
ways to authenticate, from simple passwords to OAuth and WebAuthn. This
endpoint tells a client which auth capabilities are enabled.

 */
export const authProviderList = () => {
  return fetcher<AuthProviderListOKResponse>({
    url: `/v1/auth`,
    method: "get",
  });
};

export const getAuthProviderListKey = () => [`/v1/auth`] as const;

export type AuthProviderListQueryResult = NonNullable<
  Awaited<ReturnType<typeof authProviderList>>
>;
export type AuthProviderListQueryError =
  | BadRequestResponse
  | InternalServerErrorResponse;

export const useAuthProviderList = <
  TError = BadRequestResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<
    Awaited<ReturnType<typeof authProviderList>>,
    TError
  > & { swrKey?: Key; enabled?: boolean };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getAuthProviderListKey() : null));
  const swrFn = () => authProviderList();

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
 * Register a new account with a username and password.
 */
export const authPasswordSignup = (authPasswordBody: AuthPasswordBody) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/password/signup`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: authPasswordBody,
  });
};

/**
 * Sign in to an existing account with a username and password.
 */
export const authPasswordSignin = (authPasswordBody: AuthPasswordBody) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/password/signin`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: authPasswordBody,
  });
};

/**
 * Given the requesting account does not have a password authentication,
add a password authentication method to it with the given password.

 */
export const authPasswordCreate = (
  authPasswordCreateBody: AuthPasswordCreateBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/password`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: authPasswordCreateBody,
  });
};

/**
 * Given the requesting account has a password authentication, update the
password on file.

 */
export const authPasswordUpdate = (
  authPasswordUpdateBody: AuthPasswordUpdateBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/password`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: authPasswordUpdateBody,
  });
};

/**
 * OAuth2 callback.
 */
export const oAuthProviderCallback = (
  oauthProvider: string,
  oAuthProviderCallbackBody: OAuthProviderCallbackBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/oauth/${oauthProvider}/callback`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: oAuthProviderCallbackBody,
  });
};

/**
 * Start the WebAuthn registration process by requesting a credential.

 */
export const webAuthnRequestCredential = (accountHandle: string) => {
  return fetcher<WebAuthnRequestCredentialOKResponse>({
    url: `/v1/auth/webauthn/make/${accountHandle}`,
    method: "get",
  });
};

export const getWebAuthnRequestCredentialKey = (accountHandle: string) =>
  [`/v1/auth/webauthn/make/${accountHandle}`] as const;

export type WebAuthnRequestCredentialQueryResult = NonNullable<
  Awaited<ReturnType<typeof webAuthnRequestCredential>>
>;
export type WebAuthnRequestCredentialQueryError =
  | BadRequestResponse
  | InternalServerErrorResponse;

export const useWebAuthnRequestCredential = <
  TError = BadRequestResponse | InternalServerErrorResponse,
>(
  accountHandle: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof webAuthnRequestCredential>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!accountHandle;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getWebAuthnRequestCredentialKey(accountHandle) : null));
  const swrFn = () => webAuthnRequestCredential(accountHandle);

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
 * Complete WebAuthn registration by creating a new credential.
 */
export const webAuthnMakeCredential = (
  webAuthnMakeCredentialBody: WebAuthnMakeCredentialBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/webauthn/make`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: webAuthnMakeCredentialBody,
  });
};

/**
 * Start the WebAuthn assertion for an existing account.
 */
export const webAuthnGetAssertion = (accountHandle: string) => {
  return fetcher<WebAuthnGetAssertionOKResponse>({
    url: `/v1/auth/webauthn/assert/${accountHandle}`,
    method: "get",
  });
};

export const getWebAuthnGetAssertionKey = (accountHandle: string) =>
  [`/v1/auth/webauthn/assert/${accountHandle}`] as const;

export type WebAuthnGetAssertionQueryResult = NonNullable<
  Awaited<ReturnType<typeof webAuthnGetAssertion>>
>;
export type WebAuthnGetAssertionQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useWebAuthnGetAssertion = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  accountHandle: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof webAuthnGetAssertion>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!accountHandle;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getWebAuthnGetAssertionKey(accountHandle) : null));
  const swrFn = () => webAuthnGetAssertion(accountHandle);

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
 * Complete the credential assertion and sign in to an account.
 */
export const webAuthnMakeAssertion = (
  webAuthnMakeAssertionBody: WebAuthnMakeAssertionBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/webauthn/assert`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: webAuthnMakeAssertionBody,
  });
};

/**
 * Start the authentication flow with a phone number. The handler will send
a one-time code to the provided phone number which must then be sent to
the other phone endpoint to verify the number and validate the account.

 */
export const phoneRequestCode = (
  phoneRequestCodeBody: PhoneRequestCodeBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/phone`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: phoneRequestCodeBody,
  });
};

/**
 * Complete the phone number authentication flow by submitting the one-time
code that was sent to the user's phone.

 */
export const phoneSubmitCode = (
  accountHandle: string,
  phoneSubmitCodeBody: PhoneSubmitCodeBody,
) => {
  return fetcher<AuthSuccessOKResponse>({
    url: `/v1/auth/phone/${accountHandle}`,
    method: "put",
    headers: { "Content-Type": "application/json" },
    data: phoneSubmitCodeBody,
  });
};

/**
 * Remove cookies from requesting client.
 */
export const authProviderLogout = () => {
  return fetcher<void>({ url: `/v1/auth/logout`, method: "get" });
};

export const getAuthProviderLogoutKey = () => [`/v1/auth/logout`] as const;

export type AuthProviderLogoutQueryResult = NonNullable<
  Awaited<ReturnType<typeof authProviderLogout>>
>;
export type AuthProviderLogoutQueryError =
  | BadRequestResponse
  | InternalServerErrorResponse;

export const useAuthProviderLogout = <
  TError = BadRequestResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<
    Awaited<ReturnType<typeof authProviderLogout>>,
    TError
  > & { swrKey?: Key; enabled?: boolean };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getAuthProviderLogoutKey() : null));
  const swrFn = () => authProviderLogout();

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
