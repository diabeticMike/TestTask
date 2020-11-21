package seeds

import "database/sql"

const (
	auth        = "INSERT INTO `auth` VALUES (1,'www-dfq92-sqfwf'),(2,'ffff-2918-xcas');"
	user        = "INSERT INTO `user` VALUES (1,'test'),(2,'admin'),(3,'guest');"
	authData    = "INSERT INTO `user_data` VALUES (1,'гімназія №179 міста Києва'),(2,'ліцей №227'),(3,'Медична гімназія №33 міста Києва');"
	userProfile = "INSERT INTO `user_profile` VALUES (1,'Александр','Школьный','+38050123455','ул. Сибирская 2','Киев'),(2,'Дмитрий','Арбузов','+38065133223','ул. Белая 4','Харьков'),(3,'Василий','Шпак','+38055221166','ул. Северная 5','Житомир');"
)

func RunSeeds(db *sql.DB) (err error) {
	if _, err = db.Exec(auth); err != nil {
		return err
	}
	if _, err = db.Exec(user); err != nil {
		return err
	}
	if _, err = db.Exec(authData); err != nil {
		return err
	}
	if _, err = db.Exec(userProfile); err != nil {
		return err
	}

	return nil
}
