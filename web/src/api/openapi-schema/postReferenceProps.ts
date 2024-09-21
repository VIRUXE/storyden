/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { AssetList } from "./assetList";
import type { CollectionList } from "./collectionList";
import type { LikeData } from "./likeData";
import type { Metadata } from "./metadata";
import type { PostDescription } from "./postDescription";
import type { ProfileReference } from "./profileReference";
import type { ReactList } from "./reactList";
import type { ThreadMark } from "./threadMark";
import type { ThreadTitle } from "./threadTitle";

export interface PostReferenceProps {
  assets: AssetList;
  author: ProfileReference;
  collections: CollectionList;
  description?: PostDescription;
  likes: LikeData;
  meta?: Metadata;
  reacts: ReactList;
  slug: ThreadMark;
  title: ThreadTitle;
}