-- 开始初始化数据 ;
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

INSERT INTO sys_casbin_rule VALUES (1, 'p', 'user', '/apis/v1/user-agent/*', '*', '', '', '', '', '');

INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'user2', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user2', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'user3', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user3', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'user4', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user4', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);


INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (1, 'infringe1', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (2, 'infringe2', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (3, 'infringe3', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (4, 'valuation1', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (5, 'valuation2', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL,NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (6, 'valuation3', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);

INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 3, 1, 1, 'infringement', 1, 0, '2022-12-05 21:03:55', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 4, 3, 1, 'infringement', 1, 0, '2022-12-05 21:03:55', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 3, 4, 1, 'valuation', 1, 0, '2022-12-05 21:03:55', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 4, 5, 1, 'valuation', 1, 0, '2022-12-05 21:03:55', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 5, 2, 1, 'infringement', 1, 0, '2022-12-05 21:03:55', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 5, 6, 1, 'valuation', 1, 0, '2022-12-05 21:03:55', '');

INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (1, '管理', '项目经理和文员工作', '大数据分析', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (2, '人工智能技术', '使用Pycharm并看论文', '人工智能', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (3, '前端开发', '使用VScode并学习博客', 'JavaScript', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (4, '后端开发', '使用goland并学习博客', 'go', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);

INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (1, 1, 1, '组长', '2022-10-18 18:49:24', NULL, '成功成为组长', '已完成', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (2, 1, 2, '暂无', '2022-10-18 18:49:24', NULL, '申请成为组员', '未审核', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (3, 1, 3, '暂无', '2022-10-18 18:49:24', NULL, '申请成为组长', '未审核', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (4, 2, 1, '组员', '2022-10-18 18:49:24', NULL, '申请成为组员', '未审核', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (5, 3, 1, '组员', '2022-10-18 18:49:24', NULL, '申请成为组员', '未审核', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (6, 4, 1, '组员', '2022-10-18 18:49:24', NULL, '申请成为组员', '未审核', NULL, NULL);
INSERT INTO `dept_rela` (`id`, `user_id`, `dept_id`, `mem_type`, `created_at`, `updated_at`, `mem_status`, `examine_status`, `create_by`, `update_by`) VALUES (7, 5, 2, '组员', '2022-10-18 18:49:24', NULL, '申请成为组员', '未审核', NULL, NULL);

-- 数据完成 ;