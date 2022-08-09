package conformance

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var (
	randomStrings = []string{
		"foo",
		"bar",
		"baz",
		"qux",
		"foo bar",
		"foo baz",
		"foo bar baz",
	}
)

func PopulateStruct(x interface{}) {
	populate(reflect.ValueOf(x), make(map[reflect.Type]struct{}), rand.New(rand.NewSource(time.Now().UnixMicro())), 100, .5)
}

func populate(v reflect.Value, typesSeen map[reflect.Type]struct{}, r *rand.Rand, maxCollectionSize int, shrinkFactor float32) {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(r.Int63()%2 != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(r.Int63() ^ (r.Int63() << 1))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(r.Uint64())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(r.Float64())
	case reflect.Interface:
		if v.IsNil() {
			return
		}
		populate(v.Elem(), typesSeen, r, maxCollectionSize, shrinkFactor)
	case reflect.Map:
		mapSize := 1
		if maxCollectionSize > 0 {
			mapSize = r.Intn(maxCollectionSize) + 1
		}
		mapTy := v.Type()
		newMap := reflect.MakeMapWithSize(mapTy, mapSize)
		for i := 0; i < mapSize; i++ {
			k := reflect.New(mapTy.Key()).Elem()
			populate(k, typesSeen, r, maxCollectionSize, shrinkFactor)
			v := reflect.New(mapTy.Elem()).Elem()
			populate(v, typesSeen, r, int(float32(maxCollectionSize)*shrinkFactor), shrinkFactor)
			newMap.SetMapIndex(k, v)
		}
		v.Set(newMap)
	case reflect.Slice:
		sliceSize := 1
		if maxCollectionSize > 0 {
			sliceSize = r.Intn(maxCollectionSize) + 1
		}
		sliceTy := v.Type()
		newSlice := reflect.MakeSlice(sliceTy, sliceSize, sliceSize)
		for i := 0; i < sliceSize; i++ {
			v := reflect.New(sliceTy.Elem()).Elem()
			populate(v, typesSeen, r, int(float32(maxCollectionSize)*shrinkFactor), shrinkFactor)
			newSlice.Index(i).Set(v)
		}
		v.Set(newSlice)
	case reflect.Pointer:
		if v.IsNil() {
			if !v.CanSet() {
				return
			}
			elemTy := v.Type().Elem()
			if _, ok := typesSeen[elemTy]; ok {
				return // avoid endless recursion
			}
			typesSeen[elemTy] = struct{}{}
			defer delete(typesSeen, elemTy)

			newElem := reflect.New(elemTy)
			v.Set(newElem)
		}
		populate(v.Elem(), typesSeen, r, maxCollectionSize, shrinkFactor)
	case reflect.String:
		v.SetString(randomStrings[r.Intn(len(randomStrings))])
	case reflect.Struct:
		structTy := v.Type()
		for i := 0; i < structTy.NumField(); i++ {
			field := structTy.Field(i)
			if field.IsExported() {
				populate(v.Field(i), typesSeen, r, maxCollectionSize, shrinkFactor)
			}
		}
	default:
		fmt.Println("ough")
		panic(fmt.Errorf("unhandled kind %v", v.Kind()))
	}
}
