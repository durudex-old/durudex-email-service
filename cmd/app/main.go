/*
 * Copyright © 2021-2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"

	"github.com/durudex/durudex-email-service/internal/app"
)

var configPath string // Path to config file.

// A function that is called before the main function.
func init() {
	// Add a new flag for the path to config.
	flag.StringVar(&configPath, "config", "configs/main", "Path to config file.")

	// Parsing all the flags that were specified at startup.
	flag.Parse()
}

// The main function that is called when running the application.
func main() {
	// Running a this application.
	app.Run(configPath)
}
