package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var tmt *menutable

func TestInitMenu(t *testing.T) {
	tmt = InitMenu()
	assert.NotNil(t, tmt, "InitMenu Initialized menutable, but returns nil")
	assert.NotNil(t, tmt.table, "InitMenu Initialized menutable.talbe, but returns nil")
	_, ok := tmt.table[LUNCH]
	assert.True(t, ok, "InitMenu Initialized menutable.talbe[LUNCH], but not exist")
	_, ok = tmt.table[DINNER]
	assert.True(t, ok, "InitMenu Initialized menutable.talbe[DINNER], but not exist")
}

type TestFilePath struct{}

func (TestFilePath) convertFilepath(file string) string {
	return file
}

func TestParseMenuFile(t *testing.T) {
	mockMenuTable := menutable{
		table: make(map[LunOrDin]menu),
	}
	mockMenuTable.table[LUNCH] = make(menu)
	mockMenuTable.table[DINNER] = make(menu)

	mockMenuTable.table[LUNCH][time.Monday] = append(mockMenuTable.table[LUNCH][time.Monday], []string{"돈사태메추리알조림", "백미밥*잡곡밥", "얼큰부대찌개", "칠리만두강정", "쑥갓두부무침", "콩자반", "깍두기", "식빵*딸기잼/셀프라면", "달걀후라이/셀프비빔밥", "음료2종"}...)
	mockMenuTable.table[LUNCH][time.Tuesday] = append(mockMenuTable.table[LUNCH][time.Tuesday], []string{"매콤닭볶음탕", "백미밥*잡곡밥", "열무된장국", "고구마고로케*케찹", "청포묵김가루무침", "양념깻잎지", "배추김치", "식빵*사과잼/셀프라면", "달걀후라이/셀프비빔밥", "음료2종"}...)
	mockMenuTable.table[LUNCH][time.Wednesday] = append(mockMenuTable.table[LUNCH][time.Wednesday], []string{"매콤제육불고기", "잡곡밥", "조랭이떡미역국", "모듬감자튀김*케찹", "양배추쌈*쌈장", "숙주나물무침", "깍두기", "식빵*딸기잼/셀프라면", "달걀후라이/셀프비빔밥", "음료2종"}...)
	mockMenuTable.table[LUNCH][time.Thursday] = append(mockMenuTable.table[LUNCH][time.Thursday], []string{"마파두부", "백미밥*잡곡밥", "맑은감자국", "맛초킹탕수육", "죽순부추무침", "무짠지채", "배추김치", "식빵*사과잼/셀프라면", "달걀후라이/셀프비빔밥", "음료2종"}...)
	mockMenuTable.table[LUNCH][time.Friday] = append(mockMenuTable.table[LUNCH][time.Friday], []string{"미니돈까스*수제돈까스S", "햄야채볶음밥", "맑은순두부국", "마카로니샐러드", "고춧잎무말랭이무침", "매콤콩나물무침", "배추김치", "식빵*딸기잼/셀프라면", "달걀후라이", "음료2종"}...)

	mockMenuTable.table[DINNER][time.Monday] = append(mockMenuTable.table[DINNER][time.Monday], []string{"파채너비아니구이", "백미밥*잡곡밥", "참치김치찌개", "미역줄기볶음", "감자범벅샐러드", "근대겉절이", "깍두기", "식빵*딸기잼", "달걀후라이", "음료2종"}...)
	mockMenuTable.table[DINNER][time.Tuesday] = append(mockMenuTable.table[DINNER][time.Tuesday], []string{"간장바싹불고기", "잡곡밥", "얼큰육개장", "옛날소세지전*케찹", "온두부찜*간장", "시금치무침", "배추김치", "식빵*사과잼", "달걀후라이", "음료2종"}...)
	mockMenuTable.table[DINNER][time.Wednesday] = append(mockMenuTable.table[DINNER][time.Wednesday], []string{"하이라이스", "잡곡밥", "얼큰해물짬뽕국", "순살닭강정", "푸실리파스타", "반달단무지", "깍두기", "식빵*딸기잼", "달걀후라이", "음료2종"}...)
	mockMenuTable.table[DINNER][time.Thursday] = append(mockMenuTable.table[DINNER][time.Thursday], []string{"매콤돈육떡폭찹", "잡곡밥", "얼큰우거지해장국", "매콤분모자국물떡볶이", "오징어링튀김*칠리S", "부추겉절이", "깍두기", "식빵*사과잼", "달걀후라이", "음료2종"}...)
	mockMenuTable.table[DINNER][time.Friday] = append(mockMenuTable.table[DINNER][time.Friday], []string{"간장볼어묵조림", "잡곡밥", "장칼국수", "한입부추전", "매콤마늘쫑장아찌", "후식)오렌지주스", "배추김치", "식빵*딸기잼", "달걀후라이", "음료2종"}...)

	tmt.parseMenuFile("testDiet.xlsx", TestFilePath{})

	assert.Equal(t, tmt.table, mockMenuTable.table, "tmt Equals modkMenuTable but different")
}
