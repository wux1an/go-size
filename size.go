// Package size
// Unit(Full) Unit   Bin    Value                                 Unit
// KiloByte    KB    2^10	1,024                              1,024 B
// MegaByte    MB    2^20	1,048,576                          1,024 KB
// GigaByte    GB    2^30	1,073,741,824                      1,024 MB
// TeraByte    TB    2^40	1,099,511,627,776                  1,024 GB
// PetaByte    PB    2^50	1,125,899,906,842,624              1,024 TB (not support)
// ExaByte     EB    2^60	1,152,921,504,606,846,976          1,024 PB (not support)
// ZettaByte   ZB    2^70	1,180,591,620,717,411,303,424      1,024 EB (not support)
// YottaByte   YB    2^80	1,208,925,819,614,629,174,706,176  1,024 ZB (not support)
package size

import (
	"fmt"
)

type Unit uint64

const (
	Bit  Unit = 1
	Byte      = Bit << 3
	KB        = Byte << 10
	MB        = KB << 10
	GB        = MB << 10
	TB        = GB << 10
)

func (u Unit) String() string {
	n := uint64(u)
	t := uint64(0b1111111111)
	us := []string{"PB", "TB", "GB", "MB", "KB", "Byte", "Bit"}
	vs := []uint64{
		(n & (t << 53)) >> 53, (n & (t << 43)) >> 43, (n & (t << 33)) >> 33, (n & (t << 23)) >> 23,
		(n & (t << 13)) >> 13, (n & (t << 03)) >> 3, n & 0b111,
	}

	for i := 0; i < len(us); i++ {
		if vs[i] == 0 && i != len(us)-1 {
			continue
		}

		if i == len(us)-1 || i == len(us)-2 {
			return fmt.Sprintf("%d %s", vs[i], us[i])
		} else {
			if vs[i+1] != 0 {
				return fmt.Sprintf("%d.%.0f %s", vs[i], (float64(vs[i+1])/1024)*10, us[i])
			} else {
				return fmt.Sprintf("%d %s", vs[i], us[i])
			}
		}
	}

	return "" // Unreachable code
}
