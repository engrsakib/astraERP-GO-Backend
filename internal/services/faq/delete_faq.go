package faq


func (s *FaqService) DeleteFaq(id string) error {
	
	if _, err := s.Repo.FindByID(id); err != nil {
		return err
	}

	
	return s.Repo.Delete(id)
}