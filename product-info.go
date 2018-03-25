package main

type ProductInfo struct {
	Code    string `json:"code"`
	Product struct {
		NutritionScoreDebug        string        `json:"nutrition_score_debug"`
		IngredientsFromPalmOilTags []interface{} `json:"ingredients_from_palm_oil_tags"`
		Nutriments                 struct {
			Energy               string `json:"energy"`
			SaturatedFat100G     string `json:"saturated-fat_100g"`
			SaltServing          string `json:"salt_serving"`
			Fat100G              string `json:"fat_100g"`
			FatServing           string `json:"fat_serving"`
			Carbohydrates100G    string `json:"carbohydrates_100g"`
			CarbohydratesUnit    string `json:"carbohydrates_unit"`
			Sodium               string `json:"sodium"`
			SaturatedFatValue    string `json:"saturated-fat_value"`
			Fiber100G            string `json:"fiber_100g"`
			SugarsUnit           string `json:"sugars_unit"`
			Salt                 string `json:"salt"`
			EnergyValue          string `json:"energy_value"`
			Salt100G             string `json:"salt_100g"`
			SaturatedFatUnit     string `json:"saturated-fat_unit"`
			FiberValue           string `json:"fiber_value"`
			SugarsValue          string `json:"sugars_value"`
			FiberUnit            string `json:"fiber_unit"`
			ProteinsUnit         string `json:"proteins_unit"`
			Proteins100G         string `json:"proteins_100g"`
			SodiumServing        string `json:"sodium_serving"`
			Carbohydrates        string `json:"carbohydrates"`
			CarbohydratesValue   string `json:"carbohydrates_value"`
			SaltUnit             string `json:"salt_unit"`
			SugarsServing        string `json:"sugars_serving"`
			ProteinsValue        string `json:"proteins_value"`
			Proteins             string `json:"proteins"`
			Fiber                string `json:"fiber"`
			FiberServing         string `json:"fiber_serving"`
			Fat                  string `json:"fat"`
			FatValue             string `json:"fat_value"`
			Sugars100G           string `json:"sugars_100g"`
			FatUnit              string `json:"fat_unit"`
			EnergyServing        string `json:"energy_serving"`
			Sodium100G           string `json:"sodium_100g"`
			ProteinsServing      string `json:"proteins_serving"`
			CarbohydratesServing string `json:"carbohydrates_serving"`
			Sugars               string `json:"sugars"`
			Energy100G           string `json:"energy_100g"`
			SaturatedFat         string `json:"saturated-fat"`
			EnergyUnit           string `json:"energy_unit"`
			SaltValue            string `json:"salt_value"`
			SaturatedFatServing  string `json:"saturated-fat_serving"`
		} `json:"nutriments"`
		ID                                  string        `json:"id"`
		NutritionDataPerDebugTags           []interface{} `json:"nutrition_data_per_debug_tags"`
		LastImageDatesTags                  []string      `json:"last_image_dates_tags"`
		Sortkey                             int           `json:"sortkey"`
		IngredientsText                     string        `json:"ingredients_text"`
		IngredientsThatMayBeFromPalmOilTags []interface{} `json:"ingredients_that_may_be_from_palm_oil_tags"`
		PhotographersTags                   []string      `json:"photographers_tags"`
		States                              string        `json:"states"`
		Additives                           string        `json:"additives"`
		NutrientLevels                      struct {
		} `json:"nutrient_levels"`
		IngredientsTextFr                      string        `json:"ingredients_text_fr"`
		PnnsGroups1Tags                        []string      `json:"pnns_groups_1_tags"`
		IngredientsFromOrThatMayBeFromPalmOilN int           `json:"ingredients_from_or_that_may_be_from_palm_oil_n"`
		IngredientsIdsDebug                    []string      `json:"ingredients_ids_debug"`
		IngredientsN                           string        `json:"ingredients_n"`
		Lang                                   string        `json:"lang"`
		LanguagesHierarchy                     []string      `json:"languages_hierarchy"`
		EntryDatesTags                         []string      `json:"entry_dates_tags"`
		CategoriesPrevHierarchy                []interface{} `json:"categories_prev_hierarchy"`
		NutritionDataPer                       string        `json:"nutrition_data_per"`
		PnnsGroups1                            string        `json:"pnns_groups_1"`
		LastModifiedBy                         interface{}   `json:"last_modified_by"`
		ImageIngredientsThumbURL               string        `json:"image_ingredients_thumb_url"`
		IngredientsNTags                       []string      `json:"ingredients_n_tags"`
		LastEditor                             interface{}   `json:"last_editor"`
		UnknownNutrientsTags                   []interface{} `json:"unknown_nutrients_tags"`
		ImageNutritionSmallURL                 string        `json:"image_nutrition_small_url"`
		CreatedT                               int           `json:"created_t"`
		BrandsDebugTags                        []interface{} `json:"brands_debug_tags"`
		IngredientsTextDebugTags               []interface{} `json:"ingredients_text_debug_tags"`
		ImageIngredientsURL                    string        `json:"image_ingredients_url"`
		Languages                              struct {
			EnFrench int `json:"en:french"`
		} `json:"languages"`
		UnknownIngredientsN              int           `json:"unknown_ingredients_n"`
		AdditivesPrevTags                []string      `json:"additives_prev_tags"`
		IngredientsHierarchy             []string      `json:"ingredients_hierarchy"`
		Allergens                        string        `json:"allergens"`
		EditorsTags                      []string      `json:"editors_tags"`
		AllergensHierarchy               []interface{} `json:"allergens_hierarchy"`
		IngredientsTags                  []string      `json:"ingredients_tags"`
		Creator                          string        `json:"creator"`
		StatesTags                       []string      `json:"states_tags"`
		IngredientsDebug                 []string      `json:"ingredients_debug"`
		LastEditDatesTags                []string      `json:"last_edit_dates_tags"`
		CheckersTags                     []interface{} `json:"checkers_tags"`
		ProductNameDebugTags             []interface{} `json:"product_name_debug_tags"`
		LabelsPrevHierarchy              []interface{} `json:"labels_prev_hierarchy"`
		CategoriesPrevTags               []interface{} `json:"categories_prev_tags"`
		CountriesTags                    []string      `json:"countries_tags"`
		AdditivesOldN                    int           `json:"additives_old_n"`
		Brands                           string        `json:"brands"`
		LabelsPrevTags                   []interface{} `json:"labels_prev_tags"`
		InterfaceVersionCreated          string        `json:"interface_version_created"`
		ImageNutritionURL                string        `json:"image_nutrition_url"`
		Code                             string        `json:"code"`
		Lc                               string        `json:"lc"`
		Rev                              int           `json:"rev"`
		Complete                         int           `json:"complete"`
		NutritionGrades                  string        `json:"nutrition_grades"`
		IngredientsThatMayBeFromPalmOilN int           `json:"ingredients_that_may_be_from_palm_oil_n"`
		CategoriesTags                   []interface{} `json:"categories_tags"`
		LanguagesTags                    []string      `json:"languages_tags"`
		AllergensTags                    []interface{} `json:"allergens_tags"`
		SelectedImages                   struct {
			Nutrition struct {
				Display struct {
					Fr string `json:"fr"`
				} `json:"display"`
				Thumb struct {
					Fr string `json:"fr"`
				} `json:"thumb"`
				Small struct {
					Fr string `json:"fr"`
				} `json:"small"`
			} `json:"nutrition"`
			Ingredients struct {
				Small struct {
					Fr string `json:"fr"`
				} `json:"small"`
				Display struct {
					Fr string `json:"fr"`
				} `json:"display"`
				Thumb struct {
					Fr string `json:"fr"`
				} `json:"thumb"`
			} `json:"ingredients"`
		} `json:"selected_images"`
		CodesTags   []string `json:"codes_tags"`
		Ingredients []struct {
			ID   string `json:"id"`
			Text string `json:"text"`
			Rank int    `json:"rank"`
		} `json:"ingredients"`
		CategoriesDebugTags          []interface{} `json:"categories_debug_tags"`
		AdditivesPrevN               int           `json:"additives_prev_n"`
		CountriesDebugTags           []interface{} `json:"countries_debug_tags"`
		LastImageT                   int           `json:"last_image_t"`
		LabelsDebugTags              []interface{} `json:"labels_debug_tags"`
		AdditivesOriginalTags        []string      `json:"additives_original_tags"`
		Keywords                     []string      `json:"_keywords"`
		LabelsTags                   []interface{} `json:"labels_tags"`
		AdditivesDebugTags           []interface{} `json:"additives_debug_tags"`
		ImageIngredientsSmallURL     string        `json:"image_ingredients_small_url"`
		Categories                   string        `json:"categories"`
		BrandsTags                   []string      `json:"brands_tags"`
		StatesHierarchy              []string      `json:"states_hierarchy"`
		ProductName                  string        `json:"product_name"`
		ServingQuantity              int           `json:"serving_quantity"`
		AdditivesOldTags             []interface{} `json:"additives_old_tags"`
		AdditivesN                   int           `json:"additives_n"`
		MaxImgid                     string        `json:"max_imgid"`
		IngredientsTextWithAllergens string        `json:"ingredients_text_with_allergens"`
		Images                       struct {
			Num1 struct {
				UploadedT string `json:"uploaded_t"`
				Uploader  string `json:"uploader"`
				Sizes     struct {
					Num100 struct {
						W int `json:"w"`
						H int `json:"h"`
					} `json:"100"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
			} `json:"1"`
			Num2 struct {
				Sizes struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				Uploader  string `json:"uploader"`
				UploadedT string `json:"uploaded_t"`
			} `json:"2"`
			IngredientsFr struct {
				WhiteMagic string      `json:"white_magic"`
				Y1         interface{} `json:"y1"`
				Geometry   string      `json:"geometry"`
				Angle      interface{} `json:"angle"`
				Normalize  string      `json:"normalize"`
				Rev        string      `json:"rev"`
				X2         interface{} `json:"x2"`
				X1         interface{} `json:"x1"`
				Sizes      struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num200 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"200"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						W int `json:"w"`
						H int `json:"h"`
					} `json:"full"`
				} `json:"sizes"`
				Imgid string      `json:"imgid"`
				Y2    interface{} `json:"y2"`
			} `json:"ingredients_fr"`
			NutritionFr struct {
				Y1         interface{} `json:"y1"`
				Geometry   string      `json:"geometry"`
				WhiteMagic string      `json:"white_magic"`
				Sizes      struct {
					Num100 struct {
						W int `json:"w"`
						H int `json:"h"`
					} `json:"100"`
					Num200 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"200"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				X1        interface{} `json:"x1"`
				Y2        interface{} `json:"y2"`
				Imgid     string      `json:"imgid"`
				Normalize string      `json:"normalize"`
				Angle     interface{} `json:"angle"`
				Rev       string      `json:"rev"`
				X2        interface{} `json:"x2"`
			} `json:"nutrition_fr"`
		} `json:"images"`
		PnnsGroups2              string        `json:"pnns_groups_2"`
		InterfaceVersionModified string        `json:"interface_version_modified"`
		NutrientLevelsTags       []interface{} `json:"nutrient_levels_tags"`
		ID2                      string        `json:"_id"`
		Labels                   string        `json:"labels"`
		LabelsHierarchy          []interface{} `json:"labels_hierarchy"`
		ImageNutritionThumbURL   string        `json:"image_nutrition_thumb_url"`
		CountriesHierarchy       []string      `json:"countries_hierarchy"`
		AdditivesPrev            string        `json:"additives_prev"`
		PnnsGroups2Tags          []string      `json:"pnns_groups_2_tags"`
		MiscTags                 []string      `json:"misc_tags"`
		AdditivesTags            []string      `json:"additives_tags"`
		QualityTags              []string      `json:"quality_tags"`
		InformersTags            []string      `json:"informers_tags"`
		LanguagesCodes           struct {
			Fr int `json:"fr"`
		} `json:"languages_codes"`
		UpdateKey                      string        `json:"update_key"`
		NutritionGradesTags            []string      `json:"nutrition_grades_tags"`
		NoNutritionData                string        `json:"no_nutrition_data"`
		IngredientsFromPalmOilN        int           `json:"ingredients_from_palm_oil_n"`
		LastModifiedT                  int           `json:"last_modified_t"`
		CorrectorsTags                 []string      `json:"correctors_tags"`
		IngredientsTextDebug           string        `json:"ingredients_text_debug"`
		ProductNameFr                  string        `json:"product_name_fr"`
		CategoriesHierarchy            []interface{} `json:"categories_hierarchy"`
		IngredientsTextWithAllergensFr string        `json:"ingredients_text_with_allergens_fr"`
		Countries                      string        `json:"countries"`
		AdditivesPrevOriginalTags      []string      `json:"additives_prev_original_tags"`
	} `json:"product"`
	Status        int    `json:"status"`
	StatusVerbose string `json:"status_verbose"`
}
