/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { TargetNodeSlugQueryParameter } from "./targetNodeSlugQueryParameter";

export type NodeDeleteParams = {
  /**
 * If set, child nodes will be moved to the target node. If not set, child
nodes will be moved to the root.

 */
  target_node?: TargetNodeSlugQueryParameter;
};