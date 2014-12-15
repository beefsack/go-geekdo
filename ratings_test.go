package geekdo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGraph(t *testing.T) {
	ratings, err := ParseGraph([]byte(`<div style='padding:5px;' align='center'>
	<div class='lf'>
		Ratings Breakdown: <a  href="/boardgame/50862/caligula"   >Caligula</a>
	</div>
	<div>
		<img src="//chart.googleapis.com/chart?
				cht=bvs
				&amp;chxt=x,x,y
				&amp;chxr=2,0,60.9,20
				&amp;chd=t:0,6.6,4.9,11.5,37.8,95.2,67.3,41.1,3.3,0
				&amp;chs=700x400
				&amp;chbh=r,0.3
				&amp;chco=4d89f9
				&amp;chm=t0,000000,0,0,12|t4,000000,0,1,12|t3,000000,0,2,12|t7,000000,0,3,12|t23,000000,0,4,12|t58,000000,0,5,12|t41,000000,0,6,12|t25,000000,0,7,12|t2,000000,0,8,12|t0,000000,0,9,12
				&amp;chxl=0:|1|2|3|4|5|6|7|8|9|10|1:|||||Rating|||
		"
		alt="Weight chart" /></br>
	</div>
</div>`))
	assert.NoError(t, err)
	assert.Equal(t, map[int]int{
		1:  0,
		2:  4,
		3:  3,
		4:  7,
		5:  23,
		6:  58,
		7:  41,
		8:  25,
		9:  2,
		10: 0,
	}, ratings)
}
