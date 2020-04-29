package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main()  {

	fmt.Println("hello")

	// 1.打开数据库
	db, err := bolt.Open("blockChain.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	// 将要操作数据库（改写）
    //db.Update(func(tx *bolt.Tx) error {
	//	// 2.找到抽屉bucket,没有就创建
    //	bucket := tx.Bucket([]byte("b1"))
	//	if bucket == nil {
	//		// 没有抽屉，我们需要创建
	//		bucket, err = tx.CreateBucket([]byte("b1"))
	//		if err != nil{
	//			log.Panic("创建bucket(b1)失败")
	//		}
	//	}
	//	// 3.写数据
	//	bucket.Put([]byte("11111"), []byte("hello"))
	//	bucket.Put([]byte("22222"), []byte("world"))
	//
    //	return nil
	//})

	// 4.读数据
    db.View(func(tx *bolt.Tx) error {
		// 1.找到抽屉，没有就直接报错
    	bucket := tx.Bucket([]byte("blockBucket"))
    	if bucket == nil{
    		log.Panic("读取失败")
		}
    	// 2.直接读取数据
    	v1 := bucket.Get([]byte(""))
    	v2 := bucket.Get([]byte("222222"))
    	fmt.Printf("v1 : %s\n", v1)
    	fmt.Printf("v2 : %s\n", v2)
    	return nil
	})
}
