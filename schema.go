package openbd

type (
	Audience struct {
		AudienceCode      string `json:"AudienceCode,omitempty"`
		AudienceCodeValue string `json:"AudienceCodeValue,omitempty"`
	}

	Author struct {
		Dokujikubun string  `json:"dokujikubun"`
		Listseq     float64 `json:"listseq"`
	}

	CollateralDetail struct {
		SupportingResource []SupportingResource `json:"SupportingResource,omitempty"`
		TextContent        []TextContent        `json:"TextContent,omitempty"`
	}

	Collection struct {
		CollectionSequence *CollectionSequence `json:"CollectionSequence,omitempty"`
		Collection         string              `json:"Collection,omitempty"`
		TitleDetail        struct {
			TitleElement []*TitleElement `json:"TitleElement,omitempty"`
			Title        string          `json:"Title,omitempty"`
		} `json:"TitleDetail,omitempty"`
	}

	CollectionSequence struct {
		CollectionSequenceNumber string `json:"CollectionSequenceNumber,omitempty"`
		CollectionSequence       string `json:"CollectionSequence,omitempty"`
	}

	Contributor struct {
		BiographicalNote string      `json:"BiographicalNote,omitempty"`
		ContributorRole  []string    `json:"ContributorRole,omitempty"`
		PersonName       *PersonName `json:"PersonName,omitempty"`
		SequenceNumber   string      `json:"SequenceNumber,omitempty"`
	}

	DescriptiveDetail struct {
		Audience           []Audience    `json:"Audience,omitempty"`
		Collection         *Collection   `json:"Collection,omitempty"`
		Contributor        []Contributor `json:"Contributor,omitempty"`
		Extent             []Extent      `json:"Extent,omitempty"`
		Language           []Language    `json:"Language,omitempty"`
		Measure            []Measure     `json:"Measure,omitempty"`
		ProductComposition string        `json:"ProductComposition,omitempty"`
		ProductForm        string        `json:"ProductForm,omitempty"`
		ProductFormDetail  string        `json:"ProductFormDetail,omitempty"`
		ProductPart        []ProductPart `json:"ProductPart,omitempty"`
		Subject            []Subject     `json:"Subject,omitempty"`
		TitleDetail        *TitleDetail  `json:"TitleDetail,omitempty"`
	}

	Extent struct {
		Extent      string `json:"Extent,omitempty"`
		ExtentUnit  string `json:"ExtentUnit,omitempty"`
		ExtentValue string `json:"ExtentValue,omitempty"`
	}

	Hanmoto struct {
		Author                 []Author  `json:"author,omitempty"`
		Bessoushiryou          string    `json:"bessoushiryou,omitempty"`
		Bikoutrc               string    `json:"bikoutrc,omitempty"`
		Datecreated            string    `json:"datecreated,omitempty"`
		Datejuuhanyotei        string    `json:"datejuuhanyotei,omitempty"`
		Datemodified           string    `json:"datemodified"`
		Dateshuppan            string    `json:"dateshuppan,omitempty"`
		Datezeppan             string    `json:"datezeppan,omitempty"`
		Dokushakakikomi        string    `json:"dokushakakikomi,omitempty"`
		Dokushakakikomipagesuu float64   `json:"dokushakakikomipagesuu,omitempty"`
		Furoku                 float64   `json:"furoku,omitempty"`
		Furokusonota           string    `json:"furokusonota,omitempty"`
		Gatsugougousuu         string    `json:"gatsugougousuu,omitempty"`
		Genrecodetrc           float64   `json:"genrecodetrc,omitempty"`
		Genrecodetrcjidou      float64   `json:"genrecodetrcjidou,omitempty"`
		Genshomei              string    `json:"genshomei,omitempty"`
		Han                    string    `json:"han,omitempty"`
		Hanmotokarahitokoto    string    `json:"hanmotokarahitokoto,omitempty"`
		Hastameshiyomi         bool      `json:"hastameshiyomi,omitempty"`
		Jushoujouhou           string    `json:"jushoujouhou,omitempty"`
		Jyuhan                 []Jyuhan  `json:"jyuhan,omitempty"`
		Kaisetsu105w           string    `json:"kaisetsu105w,omitempty"`
		Kankoukeitai           string    `json:"kankoukeitai,omitempty"`
		Kanrensho              string    `json:"kanrensho,omitempty"`
		Kanrenshoisbn          string    `json:"kanrenshoisbn,omitempty"`
		Kubunhanbai            bool      `json:"kubunhanbai,omitempty"`
		Lanove                 bool      `json:"lanove,omitempty"`
		Maegakinado            string    `json:"maegakinado,omitempty"`
		Ndccode                string    `json:"ndccode,omitempty"`
		Obinaiyou              string    `json:"obinaiyou,omitempty"`
		Reviews                []Reviews `json:"reviews,omitempty"`
		Rubynoumu              bool      `json:"rubynoumu,omitempty"`
		Ruishokyougousho       string    `json:"ruishokyougousho,omitempty"`
		Sonotatokkijikou       string    `json:"sonotatokkijikou,omitempty"`
		Toji                   string    `json:"toji,omitempty"`
		Tsuiki                 string    `json:"tsuiki,omitempty"`
		Zaiko                  float64   `json:"zaiko,omitempty"`
		Zasshicode             string    `json:"zasshicode,omitempty"`
	}

	Imprint struct {
		ImprintIdentifier []ImprintIdentifier `json:"ImprintIdentifier,omitempty"`
		ImprintName       string              `json:"ImprintName,omitempty"`
	}

	ImprintIdentifier struct {
		IDValue   string `json:"IDValue,omitempty"`
		ImprintID string `json:"ImprintID,omitempty"`
	}

	Jyuhan struct {
		Comment string  `json:"comment,omitempty"`
		Ctime   string  `json:"ctime,omitempty"`
		Date    string  `json:"date"`
		Suri    float64 `json:"suri,omitempty"`
	}

	Language struct {
		CountryCode  string `json:"CountryCode,omitempty"`
		LanguageCode string `json:"LanguageCode,omitempty"`
		LanguageRole string `json:"LanguageRole,omitempty"`
	}

	Measure struct {
		Measure         string `json:"Measure,omitempty"`
		MeasureUnitCode string `json:"MeasureUnitCode,omitempty"`
		Measurement     string `json:"Measurement,omitempty"`
	}

	Onix struct {
		CollateralDetail  *CollateralDetail  `json:"CollateralDetail,omitempty"`
		DescriptiveDetail *DescriptiveDetail `json:"DescriptiveDetail,omitempty"`
		Notification      string             `json:"Notification,omitempty"`
		ProductIdentifier *ProductIdentifier `json:"ProductIdentifier,omitempty"`
		ProductSupply     *ProductSupply     `json:"ProductSupply,omitempty"`
		PublishingDetail  *PublishingDetail  `json:"PublishingDetail,omitempty"`
		RecordReference   string             `json:"RecordReference,omitempty"`
	}

	PersonName struct {
		Collationkey string `json:"collationkey,omitempty"`
		Content      string `json:"content,omitempty"`
	}

	Price struct {
		CurrencyCode string      `json:"CurrencyCode,omitempty"`
		PriceAmount  string      `json:"PriceAmount,omitempty"`
		PriceDate    []PriceDate `json:"PriceDate,omitempty"`
		Price        string      `json:"Price,omitempty"`
	}

	PriceDate struct {
		Date          string `json:"Date,omitempty"`
		Price         string `json:"Price,omitempty"`
		PriceDateRole string `json:"PriceDateRole,omitempty"`
	}

	ProductIdentifier struct {
		IDValue   string `json:"IDValue,omitempty"`
		ProductID string `json:"ProductID,omitempty"`
	}

	ProductPart struct {
		NumberOfItemsOfThisForm string `json:"NumberOfItemsOfThisForm,omitempty"`
		ProductForm             string `json:"ProductForm,omitempty"`
		ProductFormDescription  string `json:"ProductFormDescription,omitempty"`
	}

	ProductSupply struct {
		SupplyDetail *SupplyDetail `json:"SupplyDetail,omitempty"`
	}

	Publisher struct {
		PublisherIdentifier []PublisherIdentifier `json:"PublisherIdentifier,omitempty"`
		PublisherName       string                `json:"PublisherName,omitempty"`
		PublishingRole      string                `json:"PublishingRole,omitempty"`
	}

	PublisherIdentifier struct {
		IDValue     string `json:"IDValue,omitempty"`
		PublisherID string `json:"PublisherID,omitempty"`
	}

	PublishingDate struct {
		Date               string `json:"Date,omitempty"`
		PublishingDateRole string `json:"PublishingDateRole,omitempty"`
	}

	PublishingDetail struct {
		Imprint        *Imprint         `json:"Imprint,omitempty"`
		Publisher      *Publisher       `json:"Publisher,omitempty"`
		PublishingDate []PublishingDate `json:"PublishingDate,omitempty"`
	}

	ResourceVersion struct {
		ResourceForm           string        `json:"ResourceForm,omitempty"`
		ResourceLink           string        `json:"ResourceLink,omitempty"`
		ResourceVersionFeature []interface{} `json:"ResourceVersionFeature,omitempty"`
	}

	ReturnsConditions struct {
		ReturnsCode string `json:"ReturnsCode,omitempty"`
	}

	Reviews struct {
		Appearance string  `json:"appearance,omitempty"`
		Choyukan   string  `json:"choyukan,omitempty"`
		Han        string  `json:"han,omitempty"`
		KubunId    float64 `json:"kubun_id,omitempty"`
		Link       string  `json:"link,omitempty"`
		PostUser   string  `json:"post_user,omitempty"`
		Reviewer   string  `json:"reviewer,omitempty"`
		Source     string  `json:"source,omitempty"`
		SourceId   float64 `json:"source_id,omitempty"`
	}

	Root struct {
		Hanmoto *Hanmoto `json:"hanmoto,omitempty"`
		Onix    *Onix    `json:"onix,omitempty"`
		Summary *Summary `json:"summary,omitempty"`
	}

	Subject struct {
		MainSubject             string `json:"MainSubject,omitempty"`
		SubjectCode             string `json:"SubjectCode,omitempty"`
		SubjectHeadingText      string `json:"SubjectHeadingText,omitempty"`
		SubjectSchemeIdentifier string `json:"SubjectSchemeIdentifier,omitempty"`
	}

	Subtitle struct {
		Collationkey string `json:"collationkey,omitempty"`
		Content      string `json:"content,omitempty"`
	}

	Summary struct {
		Author    string `json:"author,omitempty"`
		Cover     string `json:"cover,omitempty"`
		Isbn      string `json:"isbn,omitempty"`
		Pubdate   string `json:"pubdate,omitempty"`
		Publisher string `json:"publisher,omitempty"`
		Series    string `json:"series,omitempty"`
		Title     string `json:"title,omitempty"`
		Volume    string `json:"volume,omitempty"`
	}

	SupplyDetail struct {
		Price               []Price            `json:"Price,omitempty"`
		ProductAvailability string             `json:"ProductAvailability,omitempty"`
		ReturnsConditions   *ReturnsConditions `json:"ReturnsConditions,omitempty"`
	}

	SupportingResource struct {
		ContentAudience string            `json:"ContentAudience,omitempty"`
		ResourceContent string            `json:"ResourceContent,omitempty"`
		ResourceMode    string            `json:"ResourceMode,omitempty"`
		ResourceVersion []ResourceVersion `json:"ResourceVersion,omitempty"`
	}

	TextContent struct {
		ContentAudience string `json:"ContentAudience,omitempty"`
		Text            string `json:"Text,omitempty"`
		TextType        string `json:"TextType,omitempty"`
	}

	TitleDetail struct {
		TitleElement *TitleElement `json:"TitleElement,omitempty"`
		Title        string        `json:"Title,omitempty"`
	}

	TitleElement struct {
		Subtitle          *Subtitle  `json:"Subtitle,omitempty"`
		TitleElementLevel string     `json:"TitleElementLevel,omitempty"`
		TitleText         *TitleText `json:"TitleText,omitempty"`
	}

	TitleText struct {
		Collationkey string `json:"collationkey,omitempty"`
		Content      string `json:"content,omitempty"`
	}
)

func (r *Root) ISBN() string {
	return r.Summary.Isbn
}

func (r *Root) Title() string {
	return r.Summary.Title
}

func (r *Root) Author() string {
	return r.Summary.Author
}

func (r *Root) Cover() string {
	return r.Summary.Cover
}

func (r *Root) PublishedAt() string {
	return r.Summary.Pubdate
}

func (r *Root) Publisher() string {
	return r.Summary.Publisher
}

func (r *Root) Series() string {
	return r.Summary.Series
}

func (r *Root) Volume() string {
	return r.Summary.Volume
}

func (r *Root) ShortDescription() string {
	if r.Onix == nil {
		return ""
	}
	if r.Onix.CollateralDetail == nil {
		return ""
	}
	for i := range r.Onix.CollateralDetail.TextContent {
		if r.Onix.CollateralDetail.TextContent[i].TextType == "02" {
			return r.Onix.CollateralDetail.TextContent[i].Text
		}
	}
	return ""
}

func (r *Root) Description() string {
	if r.Onix == nil {
		return ""
	}
	if r.Onix.CollateralDetail == nil {
		return ""
	}
	for i := range r.Onix.CollateralDetail.TextContent {
		if r.Onix.CollateralDetail.TextContent[i].TextType == "03" {
			return r.Onix.CollateralDetail.TextContent[i].Text
		}
	}
	return ""
}

func (r *Root) JBPADescription() string {
	if r.Onix == nil {
		return ""
	}
	if r.Onix.CollateralDetail == nil {
		return ""
	}
	for i := range r.Onix.CollateralDetail.TextContent {
		if r.Onix.CollateralDetail.TextContent[i].TextType == "23" {
			return r.Onix.CollateralDetail.TextContent[i].Text
		}
	}
	return ""
}

func (r *Root) TableOfContents() string {
	if r.Onix == nil {
		return ""
	}
	if r.Onix.CollateralDetail == nil {
		return ""
	}
	for i := range r.Onix.CollateralDetail.TextContent {
		if r.Onix.CollateralDetail.TextContent[i].TextType == "04" {
			return r.Onix.CollateralDetail.TextContent[i].Text
		}
	}
	return ""
}
