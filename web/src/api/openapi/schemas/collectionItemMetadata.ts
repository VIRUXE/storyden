/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { CollectionItemMembershipType } from "./collectionItemMembershipType";
import type { ProfileReference } from "./profileReference";
import type { RelevanceScore } from "./relevanceScore";

export interface CollectionItemMetadata {
  /** The time that the item was added to the collection. */
  added_at: string;
  membership_type: CollectionItemMembershipType;
  owner: ProfileReference;
  relevance_score?: RelevanceScore;
}
