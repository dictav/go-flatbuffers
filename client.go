package main

import (
	"io/ioutil"
	"log"

	sample "github.com/dictav/go-flatbuffers/MyGame/Sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	buf, err := ioutil.ReadFile("monster.dat")
	if err != nil {
		log.Fatalln(err)
	}

	monster := sample.GetRootAsMonster(buf, 0)

	unionTable := new(flatbuffers.Table)
	if monster.Equipped(unionTable) {
		unionType := monster.EquippedType()
		if unionType == sample.EquipmentWeapon {
			// Create a `sample.Weapon` object that can be initialized with the contents
			// of the `flatbuffers.Table` (`unionTable`), which was populated by
			// `monster.Equipped()`.
			unionWeapon := new(sample.Weapon)
			unionWeapon.Init(unionTable.Bytes, unionTable.Pos)
			log.Println(string(unionWeapon.Name()))
			log.Println(unionWeapon.Damage())
		}
	}
}
