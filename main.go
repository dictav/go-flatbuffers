package main

import (
	"log"
	"net/http"

	sample "github.com/dictav/go-flatbuffers/MyGame/Sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

func createMonsterBuffer() []byte {
	builder := flatbuffers.NewBuilder(0)

	weaponOne := builder.CreateString("Sword")
	weaponTwo := builder.CreateString("Axe")

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponOne)
	sample.WeaponAddDamage(builder, 3)
	sword := sample.WeaponEnd(builder)

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponTwo)
	sample.WeaponAddDamage(builder, 5)
	axe := sample.WeaponEnd(builder)

	name := builder.CreateString("Orc")
	// Create a `vector` representing the inventory of the Orc. Each number
	// could correspond to an item that can be claimed after he is slain.
	// Note: Since we prepend the bytes, this loop iterates in reverse.
	sample.MonsterStartInventoryVector(builder, 10)
	for i := 9; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(10)

	sample.MonsterStartWeaponsVector(builder, 2)
	builder.PrependUOffsetT(axe)
	builder.PrependUOffsetT(sword)
	weapons := builder.EndVector(2)

	pos := sample.CreateVec3(builder, 1.0, 2.0, 3.0)

	sample.MonsterStart(builder)
	sample.MonsterAddPos(builder, pos)
	sample.MonsterAddHp(builder, 300)
	sample.MonsterAddName(builder, name)
	sample.MonsterAddInventory(builder, inv)
	sample.MonsterAddColor(builder, sample.ColorRed)
	sample.MonsterAddWeapons(builder, weapons)
	sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
	sample.MonsterAddEquipped(builder, axe)
	orc := sample.MonsterEnd(builder)
	builder.Finish(orc)

	return builder.FinishedBytes()
}

func monsterHandler(rw http.ResponseWriter, r *http.Request) {
	buf := createMonsterBuffer()
	rw.Write(buf)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/monster", monsterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
