package searchtree

import (
	"fmt"
	"our_antivirus/av/database"
	"strconv"
	"strings"

	"github.com/beevik/prefixtree"
)

var actualTree *prefixtree.Tree

type SignTree struct {
	name        string
	offsetBegin string
	offsetEnd   string
	dtype       string
	B           []byte
}

func (st *SignTree) Name() string {
	return strings.Join([]string{st.name, st.offsetBegin, st.dtype}, ":")
}

func (st *SignTree) Offset() (int64, error) {
	return strconv.ParseInt(st.offsetBegin, 16, 64)
}

func (st *SignTree) OffsetEnd() (int64, error) {
	return strconv.ParseInt(st.offsetEnd, 16, 64)
}

func (st *SignTree) DType() string {
	return st.dtype
}

func BuildSearchTree() error {
	tree := prefixtree.New()
	rows, err := database.GetConnection().Query("SELECT byte, offsetBegin, offsetEnd, dtype FROM signatures")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var signature []byte
		var offsetBegin string
		var offsetEnd string
		var dtype string
		err := rows.Scan(&signature, &offsetBegin, &offsetEnd, &dtype)

		if err != nil {
			fmt.Println(err)
			continue
		}

		strSign := string(signature)
		strByteHex := strings.Split(strSign, " ")

		byteSign := []byte{}

		for _, byteHexStr := range strByteHex {
			if v, err := strconv.ParseInt(byteHexStr, 16, 64); err != nil {
				fmt.Println(err)
				continue
			} else {
				byteSign = append(byteSign, byte(v))
			}
		}

		s := &SignTree{name: strSign, offsetBegin: offsetBegin, offsetEnd: offsetEnd, dtype: dtype, B: byteSign}
		tree.Add(string(byteSign), s)
	}

	actualTree = tree

	return nil
}

func GetSearchTree() *prefixtree.Tree {
	return actualTree
}
