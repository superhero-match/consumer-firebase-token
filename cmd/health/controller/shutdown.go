/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Shutdown is called and panics when consumer panics so that Health controller would not be responding and
// loadbalancer would mark consumer un-healthy and spin-up a new instance of consumer.
func (ctl *Controller) Shutdown(c *gin.Context) {
	c.Status(http.StatusOK)
	os.Exit(2)
}
