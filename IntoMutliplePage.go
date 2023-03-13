package groupietrackers

func IntoMultiplePages(NumberOfCards *int, Entry []Artists, toTurnNegative int , ForReacherchBar *[]Artists) []Cards {
	/*
	* We split the array of artist in multiple array of artist with a max of NumberOfCards
	* We enter informations for the navigation with the id of the page , the previous page and the next page , if there is a next page or not
	* Its all for the pagination in golang templates
	* toTurnNegative is used to turn the page number to negative if we are in a artist reasearch so the id will be negative
	* and with a gap of 1 in the index ( 0 will become -1 )
	* because if we not do that , the id of 0 will be missunderstood
	 */
	if *NumberOfCards == len(Entry) {
		*NumberOfCards++
	}
	var CardPagiantion []Cards
	var TmpCardsArray Cards
	TmpCardsArray.NotLastPage = true
	var TmpIndex int
	NbrPage := 0
	if toTurnNegative == -1 {
		NbrPage = 1
	}
	for index := range Entry {
		TmpIndex++
		TmpCardsArray.Array = append(TmpCardsArray.Array, Entry[index])
		TmpCardsArray.ForReacherchBar = *ForReacherchBar
		if TmpIndex == *NumberOfCards {
			TmpIndex = 0
			TmpCardsArray.PreviousPage = (NbrPage - 1) * toTurnNegative
			TmpCardsArray.NexPage = (NbrPage + 1) * toTurnNegative
			if toTurnNegative == -1 {
				TmpCardsArray.IdPage = NbrPage
			} else {
				TmpCardsArray.IdPage = NbrPage + 1
			}
			if TmpCardsArray.Array != nil {
				TmpCardsArray.IsCardIn = true
			}
			CardPagiantion = append(CardPagiantion, TmpCardsArray)
			TmpCardsArray.Array = nil
			TmpCardsArray.NotFirstPage = true
			NbrPage++
		}
	}
	TmpCardsArray.NotLastPage = false
	TmpCardsArray.PreviousPage = (NbrPage - 1) * toTurnNegative
	TmpCardsArray.IdPage = NbrPage + 1
	TmpCardsArray.NexPage = (NbrPage + 1) * toTurnNegative
	if TmpCardsArray.Array != nil {
		TmpCardsArray.IsCardIn = true
	}
	CardPagiantion = append(CardPagiantion, TmpCardsArray)
	return CardPagiantion
}
