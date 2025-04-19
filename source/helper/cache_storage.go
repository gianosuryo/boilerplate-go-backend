package helper

// func CacheStorageSaveJobRec(ctx context.Context, storage *redis.Client, jobVacancyUUID string, progressJobRec collection.CachedProgressJobRec) {
// 	progressJobRecEncoded, _ := json.Marshal(progressJobRec)
// 	key := jobVacancyUUID + ":" + os.Getenv(env.REDIS_FOLDER_JOB_REC_NAME)

// 	storage.Do(ctx, "SELECT", 0)
// 	storage.JSONSet(ctx, key, "$", progressJobRecEncoded).Err()
// 	storage.Expire(ctx, key, 24*time.Hour)
// }
