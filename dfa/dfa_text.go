package dfa

// todo 实例化过滤器

//func InitDetector(c config.Config, wordsModel model.WordsModel) *dfa.Detector {
//	detector := dfa.NewWithOptions(&dfa.Options{IgnoreCase: true,
//		Noises: []rune(" ~!@#$%^&*()_-+=?<>.—，。/\\|《》？;:：'‘；“¥·")})
//	if words, err := wordsModel.AllWords(context.Background()); err != nil {
//		panic(err.Error() + "检测器初始化失败")
//	} else {
//		for _, word := range *words {
//			detector.AddWord(word)
//		}
//	}
//
//	return detector
//}
