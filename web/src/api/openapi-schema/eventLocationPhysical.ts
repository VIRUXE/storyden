/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { EventLocationPhysicalLocationType } from "./eventLocationPhysicalLocationType";

/**
 * A physical location for an event, such as a venue, a park, a street
address, etc. This location may have a name, address, and coordinates.
A URL may also be added for a Google maps link etc.

 */
export interface EventLocationPhysical {
  address?: string;
  latitude?: number;
  location_type: EventLocationPhysicalLocationType;
  longitude?: number;
  name: string;
  url?: string;
}
