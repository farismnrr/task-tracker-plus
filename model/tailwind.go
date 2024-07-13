/** 
 * Package model provides utility functions for working with Tailwind CSS classes.
 * 
 * Variables:
 * 
 * - tailwindClasses: Variable containing a list of Tailwind CSS classes.
 *   Type: []string
 *   Description: This variable holds a list of Tailwind CSS classes used for checking if a given class represents a Tailwind CSS class.
 * 
 * Functions:
 * 
 * - RepresentsTailwind: Function to check if a list of classes represents Tailwind CSS classes.
 *   Parameters:
 *   - classes: List of classes to check.
 *     Type: []string
 *   Returns:
 *   - bool: True if the classes contain Tailwind CSS classes, otherwise false.
 *   Description: This function iterates through the list of classes and checks if any of them represent Tailwind CSS classes by comparing with the predefined list of Tailwind classes. 
 *     It returns true if at least one class represents a Tailwind CSS class, otherwise false.
 */

package model

import (
	"strings"
)

var tailwindClasses = [...]string{
	"ease",
	"duration",
	"delay",
	"transition",
	"animate",
	"bg",
	"clip",
	"color",
	"gradient",
	"from",
	"to",
	"top",
	"via",
	"border",
	"divide",
	"rounded",
	"box",
	"block",
	"hidden",
	"inline",
	"flex",
	"grid",
	"flow",
	"items",
	"content",
	"justify",
	"self",
	"order",
	"contents",
	"space",
	"focus",
	"col",
	"gap",
	"row",
	"auto",
	"place",
	"hover",
	"group",
	"h",
	"max",
	"min",
	"list",
	"m",
	"mb",
	"mr",
	"mt",
	"ml",
	"mx",
	"my",
	"lining",
	"normal",
	"oldstyle",
	"stacked",
	"opacity",
	"outline",
	"p",
	"pb",
	"pr",
	"pt",
	"pl",
	"px",
	"py",
	"align",
	"clear",
	"clearfix",
	"float",
	"inset",
	"left",
	"right",
	"bottom",
	"object",
	"z",
	"static",
	"relative",
	"absolute",
	"fixed",
	"sticky",
	"rotate",
	"skew",
	"translate",
	"class",
	"container",
	"sm:",
	"md:",
	"lg:",
	"xl:",
	"ring",
	"overscroll",
	"shadow",
	"fill",
	"stroke",
	"table",
	"antialiased",
	"subpixel",
	"text",
	"break",
	"truncate",
	"uppercase",
	"lowercase",
	"capitalize",
	"leading",
	"underline",
	"line",
	"no",
	"font",
	"italic",
	"not",
	"whitespace",
	"tracking",
	"transform",
	"origin",
	"scale",
	"diagonal",
	"cursor",
	"appearance",
	"placeholder",
	"overflow",
	"scrolling",
	"pointer",
	"resize",
	"select",
	"visible",
	"invisible",
	"sr",
	"w",
}

func RepresentsTailwind(classes []string) bool {
	for _, className := range classes {
		for _, tailwindClass := range tailwindClasses {
			if strings.HasPrefix(className, tailwindClass) {
				return true
			}
		}
	}
	return false
}
