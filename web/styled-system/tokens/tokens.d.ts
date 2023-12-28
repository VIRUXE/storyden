/* eslint-disable */
export type Token = "borders.none" | "easings.default" | "easings.linear" | "easings.in" | "easings.out" | "easings.in-out" | "durations.fastest" | "durations.faster" | "durations.fast" | "durations.normal" | "durations.slow" | "durations.slower" | "durations.slowest" | "fontWeights.thin" | "fontWeights.extralight" | "fontWeights.light" | "fontWeights.normal" | "fontWeights.medium" | "fontWeights.semibold" | "fontWeights.bold" | "fontWeights.extrabold" | "fontWeights.black" | "lineHeights.none" | "lineHeights.tight" | "lineHeights.snug" | "lineHeights.normal" | "lineHeights.relaxed" | "lineHeights.loose" | "fonts.sans" | "fonts.serif" | "fonts.mono" | "letterSpacings.tighter" | "letterSpacings.tight" | "letterSpacings.normal" | "letterSpacings.wide" | "letterSpacings.wider" | "letterSpacings.widest" | "shadows.xs" | "shadows.sm" | "shadows.md" | "shadows.lg" | "shadows.xl" | "shadows.2xl" | "shadows.inner" | "blurs.sm" | "blurs.base" | "blurs.md" | "blurs.lg" | "blurs.xl" | "blurs.2xl" | "blurs.3xl" | "spacing.0" | "spacing.1" | "spacing.2" | "spacing.3" | "spacing.4" | "spacing.5" | "spacing.6" | "spacing.7" | "spacing.8" | "spacing.9" | "spacing.10" | "spacing.11" | "spacing.12" | "spacing.14" | "spacing.16" | "spacing.20" | "spacing.24" | "spacing.28" | "spacing.32" | "spacing.36" | "spacing.40" | "spacing.44" | "spacing.48" | "spacing.52" | "spacing.56" | "spacing.60" | "spacing.64" | "spacing.72" | "spacing.80" | "spacing.96" | "spacing.0.5" | "spacing.1.5" | "spacing.2.5" | "spacing.3.5" | "sizes.0" | "sizes.1" | "sizes.2" | "sizes.3" | "sizes.4" | "sizes.5" | "sizes.6" | "sizes.7" | "sizes.8" | "sizes.9" | "sizes.10" | "sizes.11" | "sizes.12" | "sizes.14" | "sizes.16" | "sizes.20" | "sizes.24" | "sizes.28" | "sizes.32" | "sizes.36" | "sizes.40" | "sizes.44" | "sizes.48" | "sizes.52" | "sizes.56" | "sizes.60" | "sizes.64" | "sizes.72" | "sizes.80" | "sizes.96" | "sizes.0.5" | "sizes.1.5" | "sizes.2.5" | "sizes.3.5" | "sizes.xs" | "sizes.sm" | "sizes.md" | "sizes.lg" | "sizes.xl" | "sizes.2xl" | "sizes.3xl" | "sizes.4xl" | "sizes.5xl" | "sizes.6xl" | "sizes.7xl" | "sizes.8xl" | "sizes.prose" | "sizes.full" | "sizes.min" | "sizes.max" | "sizes.fit" | "sizes.breakpoint-sm" | "sizes.breakpoint-md" | "sizes.breakpoint-lg" | "sizes.breakpoint-xl" | "sizes.breakpoint-2xl" | "animations.spin" | "animations.ping" | "animations.pulse" | "animations.bounce" | "zIndex.hide" | "zIndex.base" | "zIndex.docked" | "zIndex.dropdown" | "zIndex.sticky" | "zIndex.banner" | "zIndex.overlay" | "zIndex.modal" | "zIndex.popover" | "zIndex.skipLink" | "zIndex.toast" | "zIndex.tooltip" | "radii.none" | "radii.xs" | "radii.sm" | "radii.md" | "radii.lg" | "radii.xl" | "radii.2xl" | "radii.3xl" | "radii.full" | "fontSizes.2xs" | "fontSizes.xs" | "fontSizes.5xl" | "fontSizes.6xl" | "fontSizes.7xl" | "fontSizes.8xl" | "fontSizes.9xl" | "fontSizes.sm" | "fontSizes.md" | "fontSizes.lg" | "fontSizes.xl" | "fontSizes.2xl" | "fontSizes.3xl" | "fontSizes.4xl" | "fontSizes.heading.1" | "fontSizes.heading.2" | "fontSizes.heading.3" | "fontSizes.heading.4" | "fontSizes.heading.5" | "fontSizes.heading.6" | "fontSizes.heading.variable.1" | "fontSizes.heading.variable.2" | "fontSizes.heading.variable.3" | "fontSizes.heading.variable.4" | "fontSizes.heading.variable.5" | "fontSizes.heading.variable.6" | "colors.current" | "colors.black" | "colors.white" | "colors.transparent" | "colors.rose.50" | "colors.rose.100" | "colors.rose.200" | "colors.rose.300" | "colors.rose.400" | "colors.rose.500" | "colors.rose.600" | "colors.rose.700" | "colors.rose.800" | "colors.rose.900" | "colors.rose.950" | "colors.pink.50" | "colors.pink.100" | "colors.pink.200" | "colors.pink.300" | "colors.pink.400" | "colors.pink.500" | "colors.pink.600" | "colors.pink.700" | "colors.pink.800" | "colors.pink.900" | "colors.pink.950" | "colors.fuchsia.50" | "colors.fuchsia.100" | "colors.fuchsia.200" | "colors.fuchsia.300" | "colors.fuchsia.400" | "colors.fuchsia.500" | "colors.fuchsia.600" | "colors.fuchsia.700" | "colors.fuchsia.800" | "colors.fuchsia.900" | "colors.fuchsia.950" | "colors.purple.50" | "colors.purple.100" | "colors.purple.200" | "colors.purple.300" | "colors.purple.400" | "colors.purple.500" | "colors.purple.600" | "colors.purple.700" | "colors.purple.800" | "colors.purple.900" | "colors.purple.950" | "colors.violet.50" | "colors.violet.100" | "colors.violet.200" | "colors.violet.300" | "colors.violet.400" | "colors.violet.500" | "colors.violet.600" | "colors.violet.700" | "colors.violet.800" | "colors.violet.900" | "colors.violet.950" | "colors.indigo.50" | "colors.indigo.100" | "colors.indigo.200" | "colors.indigo.300" | "colors.indigo.400" | "colors.indigo.500" | "colors.indigo.600" | "colors.indigo.700" | "colors.indigo.800" | "colors.indigo.900" | "colors.indigo.950" | "colors.blue.50" | "colors.blue.100" | "colors.blue.200" | "colors.blue.300" | "colors.blue.400" | "colors.blue.500" | "colors.blue.600" | "colors.blue.700" | "colors.blue.800" | "colors.blue.900" | "colors.blue.950" | "colors.sky.50" | "colors.sky.100" | "colors.sky.200" | "colors.sky.300" | "colors.sky.400" | "colors.sky.500" | "colors.sky.600" | "colors.sky.700" | "colors.sky.800" | "colors.sky.900" | "colors.sky.950" | "colors.cyan.50" | "colors.cyan.100" | "colors.cyan.200" | "colors.cyan.300" | "colors.cyan.400" | "colors.cyan.500" | "colors.cyan.600" | "colors.cyan.700" | "colors.cyan.800" | "colors.cyan.900" | "colors.cyan.950" | "colors.teal.50" | "colors.teal.100" | "colors.teal.200" | "colors.teal.300" | "colors.teal.400" | "colors.teal.500" | "colors.teal.600" | "colors.teal.700" | "colors.teal.800" | "colors.teal.900" | "colors.teal.950" | "colors.emerald.50" | "colors.emerald.100" | "colors.emerald.200" | "colors.emerald.300" | "colors.emerald.400" | "colors.emerald.500" | "colors.emerald.600" | "colors.emerald.700" | "colors.emerald.800" | "colors.emerald.900" | "colors.emerald.950" | "colors.green.50" | "colors.green.100" | "colors.green.200" | "colors.green.300" | "colors.green.400" | "colors.green.500" | "colors.green.600" | "colors.green.700" | "colors.green.800" | "colors.green.900" | "colors.green.950" | "colors.lime.50" | "colors.lime.100" | "colors.lime.200" | "colors.lime.300" | "colors.lime.400" | "colors.lime.500" | "colors.lime.600" | "colors.lime.700" | "colors.lime.800" | "colors.lime.900" | "colors.lime.950" | "colors.yellow.50" | "colors.yellow.100" | "colors.yellow.200" | "colors.yellow.300" | "colors.yellow.400" | "colors.yellow.500" | "colors.yellow.600" | "colors.yellow.700" | "colors.yellow.800" | "colors.yellow.900" | "colors.yellow.950" | "colors.amber.50" | "colors.amber.100" | "colors.amber.200" | "colors.amber.300" | "colors.amber.400" | "colors.amber.500" | "colors.amber.600" | "colors.amber.700" | "colors.amber.800" | "colors.amber.900" | "colors.amber.950" | "colors.orange.50" | "colors.orange.100" | "colors.orange.200" | "colors.orange.300" | "colors.orange.400" | "colors.orange.500" | "colors.orange.600" | "colors.orange.700" | "colors.orange.800" | "colors.orange.900" | "colors.orange.950" | "colors.red.50" | "colors.red.100" | "colors.red.200" | "colors.red.300" | "colors.red.400" | "colors.red.500" | "colors.red.600" | "colors.red.700" | "colors.red.800" | "colors.red.900" | "colors.red.950" | "colors.neutral.50" | "colors.neutral.100" | "colors.neutral.200" | "colors.neutral.300" | "colors.neutral.400" | "colors.neutral.500" | "colors.neutral.600" | "colors.neutral.700" | "colors.neutral.800" | "colors.neutral.900" | "colors.neutral.950" | "colors.stone.50" | "colors.stone.100" | "colors.stone.200" | "colors.stone.300" | "colors.stone.400" | "colors.stone.500" | "colors.stone.600" | "colors.stone.700" | "colors.stone.800" | "colors.stone.900" | "colors.stone.950" | "colors.zinc.50" | "colors.zinc.100" | "colors.zinc.200" | "colors.zinc.300" | "colors.zinc.400" | "colors.zinc.500" | "colors.zinc.600" | "colors.zinc.700" | "colors.zinc.800" | "colors.zinc.900" | "colors.zinc.950" | "colors.gray.50" | "colors.gray.100" | "colors.gray.200" | "colors.gray.300" | "colors.gray.400" | "colors.gray.500" | "colors.gray.600" | "colors.gray.700" | "colors.gray.800" | "colors.gray.900" | "colors.gray.950" | "colors.slate.50" | "colors.slate.100" | "colors.slate.200" | "colors.slate.300" | "colors.slate.400" | "colors.slate.500" | "colors.slate.600" | "colors.slate.700" | "colors.slate.800" | "colors.slate.900" | "colors.slate.950" | "colors.whiteAlpha.50" | "colors.whiteAlpha.100" | "colors.whiteAlpha.200" | "colors.whiteAlpha.300" | "colors.whiteAlpha.400" | "colors.whiteAlpha.500" | "colors.whiteAlpha.600" | "colors.whiteAlpha.700" | "colors.whiteAlpha.800" | "colors.whiteAlpha.900" | "colors.blackAlpha.50" | "colors.blackAlpha.100" | "colors.blackAlpha.200" | "colors.blackAlpha.300" | "colors.blackAlpha.400" | "colors.blackAlpha.500" | "colors.blackAlpha.600" | "colors.blackAlpha.700" | "colors.blackAlpha.800" | "colors.blackAlpha.900" | "breakpoints.sm" | "breakpoints.md" | "breakpoints.lg" | "breakpoints.xl" | "breakpoints.2xl" | "blurs.frosted" | "colors.accent.50" | "colors.accent.100" | "colors.accent.200" | "colors.accent.300" | "colors.accent.400" | "colors.accent.500" | "colors.accent.600" | "colors.accent.700" | "colors.accent.800" | "colors.accent.900" | "colors.accent" | "colors.accent.text.50" | "colors.accent.text.100" | "colors.accent.text.200" | "colors.accent.text.300" | "colors.accent.text.400" | "colors.accent.text.500" | "colors.accent.text.600" | "colors.accent.text.700" | "colors.accent.text.800" | "colors.accent.text.900" | "colors.accent.text" | "colors.accent.dark.50" | "colors.accent.dark.100" | "colors.accent.dark.200" | "colors.accent.dark.300" | "colors.accent.dark.400" | "colors.accent.dark.500" | "colors.accent.dark.600" | "colors.accent.dark.700" | "colors.accent.dark.800" | "colors.accent.dark.900" | "colors.accent.dark" | "colors.accent.dark.text.50" | "colors.accent.dark.text.100" | "colors.accent.dark.text.200" | "colors.accent.dark.text.300" | "colors.accent.dark.text.400" | "colors.accent.dark.text.500" | "colors.accent.dark.text.600" | "colors.accent.dark.text.700" | "colors.accent.dark.text.800" | "colors.accent.dark.text.900" | "colors.accent.dark.text" | "colors.bg.canvas" | "colors.bg.default" | "colors.bg.subtle" | "colors.bg.muted" | "colors.bg.emphasized" | "colors.bg.disabled" | "colors.bg.destructive" | "colors.fg.default" | "colors.fg.muted" | "colors.fg.subtle" | "colors.fg.disabled" | "colors.fg.destructive" | "colors.border.default" | "colors.border.muted" | "colors.border.subtle" | "colors.border.disabled" | "colors.border.outline" | "colors.border.accent" | "colors.conicGradient" | "spacing.safeBottom" | "spacing.-1" | "spacing.-2" | "spacing.-3" | "spacing.-4" | "spacing.-5" | "spacing.-6" | "spacing.-7" | "spacing.-8" | "spacing.-9" | "spacing.-10" | "spacing.-11" | "spacing.-12" | "spacing.-14" | "spacing.-16" | "spacing.-20" | "spacing.-24" | "spacing.-28" | "spacing.-32" | "spacing.-36" | "spacing.-40" | "spacing.-44" | "spacing.-48" | "spacing.-52" | "spacing.-56" | "spacing.-60" | "spacing.-64" | "spacing.-72" | "spacing.-80" | "spacing.-96" | "spacing.-0.5" | "spacing.-1.5" | "spacing.-2.5" | "spacing.-3.5" | "spacing.-safeBottom" | "colors.colorPalette.50" | "colors.colorPalette.100" | "colors.colorPalette.200" | "colors.colorPalette.300" | "colors.colorPalette.400" | "colors.colorPalette.500" | "colors.colorPalette.600" | "colors.colorPalette.700" | "colors.colorPalette.800" | "colors.colorPalette.900" | "colors.colorPalette.950" | "colors.colorPalette.text.50" | "colors.colorPalette.text.100" | "colors.colorPalette.text.200" | "colors.colorPalette.text.300" | "colors.colorPalette.text.400" | "colors.colorPalette.text.500" | "colors.colorPalette.text.600" | "colors.colorPalette.text.700" | "colors.colorPalette.text.800" | "colors.colorPalette.text.900" | "colors.colorPalette.text" | "colors.colorPalette.dark.50" | "colors.colorPalette.dark.100" | "colors.colorPalette.dark.200" | "colors.colorPalette.dark.300" | "colors.colorPalette.dark.400" | "colors.colorPalette.dark.500" | "colors.colorPalette.dark.600" | "colors.colorPalette.dark.700" | "colors.colorPalette.dark.800" | "colors.colorPalette.dark.900" | "colors.colorPalette.dark" | "colors.colorPalette.dark.text.50" | "colors.colorPalette.dark.text.100" | "colors.colorPalette.dark.text.200" | "colors.colorPalette.dark.text.300" | "colors.colorPalette.dark.text.400" | "colors.colorPalette.dark.text.500" | "colors.colorPalette.dark.text.600" | "colors.colorPalette.dark.text.700" | "colors.colorPalette.dark.text.800" | "colors.colorPalette.dark.text.900" | "colors.colorPalette.dark.text" | "colors.colorPalette.canvas" | "colors.colorPalette.default" | "colors.colorPalette.subtle" | "colors.colorPalette.muted" | "colors.colorPalette.emphasized" | "colors.colorPalette.disabled" | "colors.colorPalette.destructive" | "colors.colorPalette.outline" | "colors.colorPalette.accent"

export type ColorPalette = "rose" | "pink" | "fuchsia" | "purple" | "violet" | "indigo" | "blue" | "sky" | "cyan" | "teal" | "emerald" | "green" | "lime" | "yellow" | "amber" | "orange" | "red" | "neutral" | "stone" | "zinc" | "gray" | "slate" | "whiteAlpha" | "blackAlpha" | "accent" | "accent.text" | "accent.dark" | "accent.dark.text" | "bg" | "fg" | "border"

export type BorderToken = "none"

export type EasingToken = "default" | "linear" | "in" | "out" | "in-out"

export type DurationToken = "fastest" | "faster" | "fast" | "normal" | "slow" | "slower" | "slowest"

export type FontWeightToken = "thin" | "extralight" | "light" | "normal" | "medium" | "semibold" | "bold" | "extrabold" | "black"

export type LineHeightToken = "none" | "tight" | "snug" | "normal" | "relaxed" | "loose"

export type FontToken = "sans" | "serif" | "mono"

export type LetterSpacingToken = "tighter" | "tight" | "normal" | "wide" | "wider" | "widest"

export type ShadowToken = "xs" | "sm" | "md" | "lg" | "xl" | "2xl" | "inner"

export type BlurToken = "sm" | "base" | "md" | "lg" | "xl" | "2xl" | "3xl" | "frosted"

export type SpacingToken = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" | "10" | "11" | "12" | "14" | "16" | "20" | "24" | "28" | "32" | "36" | "40" | "44" | "48" | "52" | "56" | "60" | "64" | "72" | "80" | "96" | "0.5" | "1.5" | "2.5" | "3.5" | "safeBottom" | "-1" | "-2" | "-3" | "-4" | "-5" | "-6" | "-7" | "-8" | "-9" | "-10" | "-11" | "-12" | "-14" | "-16" | "-20" | "-24" | "-28" | "-32" | "-36" | "-40" | "-44" | "-48" | "-52" | "-56" | "-60" | "-64" | "-72" | "-80" | "-96" | "-0.5" | "-1.5" | "-2.5" | "-3.5" | "-safeBottom"

export type SizeToken = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" | "10" | "11" | "12" | "14" | "16" | "20" | "24" | "28" | "32" | "36" | "40" | "44" | "48" | "52" | "56" | "60" | "64" | "72" | "80" | "96" | "0.5" | "1.5" | "2.5" | "3.5" | "xs" | "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "4xl" | "5xl" | "6xl" | "7xl" | "8xl" | "prose" | "full" | "min" | "max" | "fit" | "breakpoint-sm" | "breakpoint-md" | "breakpoint-lg" | "breakpoint-xl" | "breakpoint-2xl"

export type AnimationToken = "spin" | "ping" | "pulse" | "bounce"

export type ZIndexToken = "hide" | "base" | "docked" | "dropdown" | "sticky" | "banner" | "overlay" | "modal" | "popover" | "skipLink" | "toast" | "tooltip"

export type RadiusToken = "none" | "xs" | "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "full"

export type FontSizeToken = "2xs" | "xs" | "5xl" | "6xl" | "7xl" | "8xl" | "9xl" | "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "4xl" | "heading.1" | "heading.2" | "heading.3" | "heading.4" | "heading.5" | "heading.6" | "heading.variable.1" | "heading.variable.2" | "heading.variable.3" | "heading.variable.4" | "heading.variable.5" | "heading.variable.6"

export type ColorToken = "current" | "black" | "white" | "transparent" | "rose.50" | "rose.100" | "rose.200" | "rose.300" | "rose.400" | "rose.500" | "rose.600" | "rose.700" | "rose.800" | "rose.900" | "rose.950" | "pink.50" | "pink.100" | "pink.200" | "pink.300" | "pink.400" | "pink.500" | "pink.600" | "pink.700" | "pink.800" | "pink.900" | "pink.950" | "fuchsia.50" | "fuchsia.100" | "fuchsia.200" | "fuchsia.300" | "fuchsia.400" | "fuchsia.500" | "fuchsia.600" | "fuchsia.700" | "fuchsia.800" | "fuchsia.900" | "fuchsia.950" | "purple.50" | "purple.100" | "purple.200" | "purple.300" | "purple.400" | "purple.500" | "purple.600" | "purple.700" | "purple.800" | "purple.900" | "purple.950" | "violet.50" | "violet.100" | "violet.200" | "violet.300" | "violet.400" | "violet.500" | "violet.600" | "violet.700" | "violet.800" | "violet.900" | "violet.950" | "indigo.50" | "indigo.100" | "indigo.200" | "indigo.300" | "indigo.400" | "indigo.500" | "indigo.600" | "indigo.700" | "indigo.800" | "indigo.900" | "indigo.950" | "blue.50" | "blue.100" | "blue.200" | "blue.300" | "blue.400" | "blue.500" | "blue.600" | "blue.700" | "blue.800" | "blue.900" | "blue.950" | "sky.50" | "sky.100" | "sky.200" | "sky.300" | "sky.400" | "sky.500" | "sky.600" | "sky.700" | "sky.800" | "sky.900" | "sky.950" | "cyan.50" | "cyan.100" | "cyan.200" | "cyan.300" | "cyan.400" | "cyan.500" | "cyan.600" | "cyan.700" | "cyan.800" | "cyan.900" | "cyan.950" | "teal.50" | "teal.100" | "teal.200" | "teal.300" | "teal.400" | "teal.500" | "teal.600" | "teal.700" | "teal.800" | "teal.900" | "teal.950" | "emerald.50" | "emerald.100" | "emerald.200" | "emerald.300" | "emerald.400" | "emerald.500" | "emerald.600" | "emerald.700" | "emerald.800" | "emerald.900" | "emerald.950" | "green.50" | "green.100" | "green.200" | "green.300" | "green.400" | "green.500" | "green.600" | "green.700" | "green.800" | "green.900" | "green.950" | "lime.50" | "lime.100" | "lime.200" | "lime.300" | "lime.400" | "lime.500" | "lime.600" | "lime.700" | "lime.800" | "lime.900" | "lime.950" | "yellow.50" | "yellow.100" | "yellow.200" | "yellow.300" | "yellow.400" | "yellow.500" | "yellow.600" | "yellow.700" | "yellow.800" | "yellow.900" | "yellow.950" | "amber.50" | "amber.100" | "amber.200" | "amber.300" | "amber.400" | "amber.500" | "amber.600" | "amber.700" | "amber.800" | "amber.900" | "amber.950" | "orange.50" | "orange.100" | "orange.200" | "orange.300" | "orange.400" | "orange.500" | "orange.600" | "orange.700" | "orange.800" | "orange.900" | "orange.950" | "red.50" | "red.100" | "red.200" | "red.300" | "red.400" | "red.500" | "red.600" | "red.700" | "red.800" | "red.900" | "red.950" | "neutral.50" | "neutral.100" | "neutral.200" | "neutral.300" | "neutral.400" | "neutral.500" | "neutral.600" | "neutral.700" | "neutral.800" | "neutral.900" | "neutral.950" | "stone.50" | "stone.100" | "stone.200" | "stone.300" | "stone.400" | "stone.500" | "stone.600" | "stone.700" | "stone.800" | "stone.900" | "stone.950" | "zinc.50" | "zinc.100" | "zinc.200" | "zinc.300" | "zinc.400" | "zinc.500" | "zinc.600" | "zinc.700" | "zinc.800" | "zinc.900" | "zinc.950" | "gray.50" | "gray.100" | "gray.200" | "gray.300" | "gray.400" | "gray.500" | "gray.600" | "gray.700" | "gray.800" | "gray.900" | "gray.950" | "slate.50" | "slate.100" | "slate.200" | "slate.300" | "slate.400" | "slate.500" | "slate.600" | "slate.700" | "slate.800" | "slate.900" | "slate.950" | "whiteAlpha.50" | "whiteAlpha.100" | "whiteAlpha.200" | "whiteAlpha.300" | "whiteAlpha.400" | "whiteAlpha.500" | "whiteAlpha.600" | "whiteAlpha.700" | "whiteAlpha.800" | "whiteAlpha.900" | "blackAlpha.50" | "blackAlpha.100" | "blackAlpha.200" | "blackAlpha.300" | "blackAlpha.400" | "blackAlpha.500" | "blackAlpha.600" | "blackAlpha.700" | "blackAlpha.800" | "blackAlpha.900" | "accent.50" | "accent.100" | "accent.200" | "accent.300" | "accent.400" | "accent.500" | "accent.600" | "accent.700" | "accent.800" | "accent.900" | "accent" | "accent.text.50" | "accent.text.100" | "accent.text.200" | "accent.text.300" | "accent.text.400" | "accent.text.500" | "accent.text.600" | "accent.text.700" | "accent.text.800" | "accent.text.900" | "accent.text" | "accent.dark.50" | "accent.dark.100" | "accent.dark.200" | "accent.dark.300" | "accent.dark.400" | "accent.dark.500" | "accent.dark.600" | "accent.dark.700" | "accent.dark.800" | "accent.dark.900" | "accent.dark" | "accent.dark.text.50" | "accent.dark.text.100" | "accent.dark.text.200" | "accent.dark.text.300" | "accent.dark.text.400" | "accent.dark.text.500" | "accent.dark.text.600" | "accent.dark.text.700" | "accent.dark.text.800" | "accent.dark.text.900" | "accent.dark.text" | "bg.canvas" | "bg.default" | "bg.subtle" | "bg.muted" | "bg.emphasized" | "bg.disabled" | "bg.destructive" | "fg.default" | "fg.muted" | "fg.subtle" | "fg.disabled" | "fg.destructive" | "border.default" | "border.muted" | "border.subtle" | "border.disabled" | "border.outline" | "border.accent" | "conicGradient" | "colorPalette.50" | "colorPalette.100" | "colorPalette.200" | "colorPalette.300" | "colorPalette.400" | "colorPalette.500" | "colorPalette.600" | "colorPalette.700" | "colorPalette.800" | "colorPalette.900" | "colorPalette.950" | "colorPalette.text.50" | "colorPalette.text.100" | "colorPalette.text.200" | "colorPalette.text.300" | "colorPalette.text.400" | "colorPalette.text.500" | "colorPalette.text.600" | "colorPalette.text.700" | "colorPalette.text.800" | "colorPalette.text.900" | "colorPalette.text" | "colorPalette.dark.50" | "colorPalette.dark.100" | "colorPalette.dark.200" | "colorPalette.dark.300" | "colorPalette.dark.400" | "colorPalette.dark.500" | "colorPalette.dark.600" | "colorPalette.dark.700" | "colorPalette.dark.800" | "colorPalette.dark.900" | "colorPalette.dark" | "colorPalette.dark.text.50" | "colorPalette.dark.text.100" | "colorPalette.dark.text.200" | "colorPalette.dark.text.300" | "colorPalette.dark.text.400" | "colorPalette.dark.text.500" | "colorPalette.dark.text.600" | "colorPalette.dark.text.700" | "colorPalette.dark.text.800" | "colorPalette.dark.text.900" | "colorPalette.dark.text" | "colorPalette.canvas" | "colorPalette.default" | "colorPalette.subtle" | "colorPalette.muted" | "colorPalette.emphasized" | "colorPalette.disabled" | "colorPalette.destructive" | "colorPalette.outline" | "colorPalette.accent"

export type BreakpointToken = "sm" | "md" | "lg" | "xl" | "2xl"

export type AnimationName = "shimmer"

export type Tokens = {
		borders: BorderToken
		easings: EasingToken
		durations: DurationToken
		fontWeights: FontWeightToken
		lineHeights: LineHeightToken
		fonts: FontToken
		letterSpacings: LetterSpacingToken
		shadows: ShadowToken
		blurs: BlurToken
		spacing: SpacingToken
		sizes: SizeToken
		animations: AnimationToken
		zIndex: ZIndexToken
		radii: RadiusToken
		fontSizes: FontSizeToken
		colors: ColorToken
		breakpoints: BreakpointToken
		animationName: AnimationName
} & { [token: string]: never }

export type TokenCategory = "zIndex" | "opacity" | "colors" | "fonts" | "fontSizes" | "fontWeights" | "lineHeights" | "letterSpacings" | "sizes" | "shadows" | "spacing" | "radii" | "borders" | "durations" | "easings" | "animations" | "blurs" | "gradients" | "breakpoints" | "assets"