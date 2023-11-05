package main

import "fmt"

type Scene struct {
	LeftLocationId  int
	RightLocationId int
	Description     string
	Win             bool
}

func (s *Scene) getNextLocation(walkLeft bool) int {
	if walkLeft {
		return s.LeftLocationId
	}
	return s.RightLocationId
}

func NewScene(d string, Left int, Right int, Win bool) Scene {
	return Scene{LeftLocationId: Left, RightLocationId: Right, Description: d, Win: Win}
}

func (s *Scene) MoveNext(walkLeft bool) (*Scene, error) {
	id := s.getNextLocation(walkLeft)
	if id == 0 {
		return nil,
			fmt.Errorf("hero dead")
	}

	Current, exist := locat[id]
	if !exist {
		return nil,
			fmt.Errorf("hero dead")
	}
	return &Current, nil
}

var locat = (map[int]Scene{
	1: NewScene("First location: Stephen opened his eyes near the cave and in front of him there were two paths: to the left into the path, and to the right the entrance to the cave", 2, 3, false), // айди передать как по листу + добавить описание. вибрать ласт сцену где он вин добавить еще 1 парметр вин или не вин
	2: NewScene("Second Location: Stephen walked along a path near the forest and saw an animal", 6, 7, false),                                                                                       // дописать дискрипшн
	3: NewScene("Third Location: Stephen went into a cave where it was cold and dark ", 7, 11, false),                                                                                                //
	4: NewScene("Fourth Location : Stephen found his way home, the hero won ", 0, 0, true),                                                                                                           //win локация
	5: NewScene("Fifth Location: Stephen walked further and saw a safe on the left, and a path further on the right", 0, 4, false),                                                                   // поменять
	6: NewScene("Sixth Location: Stephen found himself in a deserted camp and he has a choice: on the left is a bed on which he can lie down and rest, and on the right is a path to move on.", 5, 0, false),
	7: NewScene("Seven Location: The hero fell into a trap. Choose which hand he can help himself with: with his left hand he simply tries to break the trap, and with his right hand he tries to get a knife and cut the trap", 0, 8, false),
	8: NewScene("Eight Location: Stephen broke free and got on the right road", 11, 4, false),
})

func main() {

	first := locat[1]
	Current := &first
	//fmt.Println("Stiven rodilsia")
	for {
		fmt.Println(Current.Description, "choose 1 or 2 , left = 1 , right = 2")
		var UserInput int
		fmt.Scan(&UserInput)
		fmt.Println("User input : ", UserInput)
		walkLeft := UserInput == 1

		Next, err := Current.MoveNext(walkLeft)
		if err != nil {
			fmt.Println("Next is nill = finish", err.Error())
			break
		}
		//тут нужна проверка на Win локацию если вин то написать
		if Next.Win {
			fmt.Println(Next.Description)
			break
		}

		Current = Next
	}

}
