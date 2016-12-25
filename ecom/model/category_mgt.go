package model

type ICategoryRepo interface {
	Read(id string) *Category
	All() []Category
}

type CategoryMgt struct {
	Repo ICategoryRepo
}

func (me *CategoryMgt) Read(id string) *Category {
	return me.Repo.Read(id)
}

func buildTreeInt(categories []*Category, node *Category) *Category {
	if node == nil {
		var root *Category
		nodefound := false
		// find root
		for _, category := range categories {
			if category.ParentId == "" {
				root = category
				nodefound = true
				break
			}
		}

		if !nodefound {
			return nil
		}

		buildTreeInt(categories, root)
		return root
	}
	
	for _, category := range categories {
		if category.ParentId == node.Id {
			node.Children = append(node.Children, category)
			buildTreeInt(categories, category)
		}
	}
	return nil
}

func BuildTree(categories []*Category) *Category {
	return buildTreeInt(categories, nil)
}

func (me *CategoryMgt) GetRoot() *Category {
	categories := me.Repo.All()
	categoryps := make([]*Category, len(categories))
	for idx, cat := range categories {
		categoryps[idx] = &cat
	}
	return BuildTree(categoryps)
}
