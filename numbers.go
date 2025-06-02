// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

// TruncateList removes duplicates fro a list.
func TruncateList(in []int, limit int) []int {
	var out []int

	i := 0

	for i < len(in)-1 {
		if in[i+1] != in[i] {
			out = append(out, in[i])
			i++

			continue
		}

		tmp := make([]int, 0)

		for _, d := range in[i:] {
			if d == in[i] {
				tmp = append(tmp, d)
			}
		}

		if len(tmp) <= limit {
			out = append(out, tmp...)
		}
		i += len(tmp)
	}

	if i < len(in) {
		out = append(out, in[i])
	}

	return out
}
