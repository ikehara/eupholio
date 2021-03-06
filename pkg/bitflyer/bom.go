/*
 * Eupholio - A portfolio tracker tool for cryptocurrency
 * Copyright (C) 2021 Kiyoshi Nakao
 *
 * This file is part of Eupholio.
 *
 * Eupholio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Eupholio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Eupholio.  If not, see <http://www.gnu.org/licenses/>.
 */

package bitflyer

import (
	"bufio"
	"io"
)

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

func NewReader(r io.Reader) io.Reader {
	buf := bufio.NewReader(r)
	b, err := buf.Peek(3)
	if err != nil {
		return buf
	}
	if b[0] == bom0 && b[1] == bom1 && b[2] == bom2 {
		buf.Discard(3)
	}
	return buf
}
