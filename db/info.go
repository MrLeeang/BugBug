package db

import "BugBug/utils"

func Test(postid interface{}) []map[string]string {
	sql_1 := "select * from `fb_posts` where `fb_posts`.`id` = ? and `fb_posts`.`deleted_at` is null limit 1"
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}




func Test2(postid interface{}) []map[string]string {
	sql_1 := "select `fb_posts`.*, (select count(*) from `fb_votes` where `fb_posts`.`id` = `fb_votes`.`pid`) as `votes_count`, (select count(*) from `fb_post_comments` where `fb_posts`.`id` = `fb_post_comments`.`pid`) as `comments_count`, (select count(*) from `fb_adopts` where `fb_posts`.`id` = `fb_adopts`.`pid`) as `adopts_count` from `fb_posts` where `fb_posts`.`id` = ? and `fb_posts`.`deleted_at` is null limit 1"
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}



func User(postid interface{}) []map[string]string {
	sql_1 := "select `id`, `nickname`, `avatar`, `sex` from `fb_users` where `fb_users`.`id` in (select t.uid from(select * from `fb_posts` where `fb_posts`.`id` = ? and `fb_posts`.`deleted_at` is null limit 1 )as t) and `fb_users`.`deleted_at` is null   "
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}

func Circle(postid interface{}) []map[string]string{
	sql_1 := " select `id`, `name`, `logo` from `fb_circles` where `fb_circles`.`id` in (select t.cid from(select * from `fb_posts` where `fb_posts`.`id` = ? and `fb_posts`.`deleted_at` is null limit 1 )as t) and `fb_circles`.`deleted_at` is null "
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}



func PostIms(postid interface{}) []map[string]string{
	sql_1 := " select * from `fb_post_imgs` where `fb_post_imgs`.`pid` in (?) and `fb_post_imgs`.`deleted_at` is null"
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}


func PostVideo(postid interface{}) []map[string]string{
	sql_1 := "select * from `fb_post_videos` where `fb_post_videos`.`pid` in (?) and `fb_post_videos`.`deleted_at` is null "
	// fmt.Println(sql_1)
	results, err := Engine.SQL(sql_1, postid).QueryString()
	if err != nil {
		utils.UtilsLogger.Error(err)
	}
	return results
}
