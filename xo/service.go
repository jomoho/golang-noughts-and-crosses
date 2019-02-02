package xo

var games []Game
var tokens = []string{"X", "O"}

func makeField() string{
	return "***\n***\n***\n"
}

func makeFieldMove(field string, token string, x int, y int) string{
	out := []rune(field)
	i := x + 4 * y

	out[i] = []rune(token)[0]
	return string(out)
}

func validPos(x int, y int) bool{
	return x >= 0 && x < 3 && y >= 0 && y < 3
}

func validMove(field string, x int, y int) bool{
	if ! validPos(x,y){
		return false
	}
	i := x + 4 * y
	return field[i] == '*'
}

func detectWinner(field string) string {
	rfield := []rune(field)

	//vertical
	for x:= 0; x < 3; x++ {
		last := rfield[x]
		sum := 0
		for y:= 0; y < 3; y++ {
			i := x + 4 * y
			if rfield[i] == last {
				sum++
			}
		}
		if(sum == 3 && last != '*'){
			return string(last)
		}
	}
	//horizontal
	for y:= 0; y < 3; y++ {
		last := rfield[y*4]
		sum := 0
		for x:= 0; x < 3; x++ {
			i := x + 4 * y
			if rfield[i] == last {
				sum++
			}
		}
		if(sum == 3 && last != '*'){
			return string(last)
		}
	}

	//diagonal \
	last := rfield[0]
	sum := 0
	for x:= 0; x < 3; x++ {
		y := x
		i := x + 4 * y
		if rfield[i] == last {
			sum++
		}
	}
	if(sum == 3 && last != '*'){
		return string(last)
	}

	//diagonal /
	last = rfield[0 + 4 * 2]
	sum = 0
	for x:= 0; x < 3; x++ {
		y := 2-x
		i := x + 4 * y
		if rfield[i] == last {
			sum++
		}
	}
	if(sum == 3 && last != '*'){
		return string(last)
	}

	return "none"
}