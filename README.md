# KutoUI/Icons

[Lucide Icons](https://lucide.dev) packaged up to be used as 
[templ](https://templ.guide/) components.

### Usage:

See the list of icons [here](https://lucide.dev/icons/). Take 
for example the `chevron-down` icon, to use it convert the name 
to capital first letter CamelCase. So in this instance it would 
be: `@icons.ChevronDown()`.

### Options:

Each icon takes an `icons.IconOptions{}` struct for configuration 
with the following options:
| Name        | Type      | Clamp            | Default      |
| ----------- | --------- | ---------------- | ------------ |
| Color       | `string`  |                  | currentColor |
| StrokeWidth | `float32` | clamped 0.5 <> 3 | 2            |
| Size        | `int`     | clamped 16 <> 48 | 24           |