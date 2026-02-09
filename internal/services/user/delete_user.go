package user



// DeleteUser Service
func (s *UserService) DeleteUser(id string) error {
    
    _, err := s.Repo.GetUserByID(id)
    if err != nil {
        return err
    }
    
    
    return s.Repo.DeleteUser(id)
}