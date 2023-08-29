/*
 Navicat Premium Data Transfer

 Source Server         : 本地pgsql
 Source Server Type    : PostgreSQL
 Source Server Version : 150004 (150004)
 Source Host           : localhost:15432
 Source Catalog        : devserve
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150004 (150004)
 File Encoding         : 65001

 Date: 13/08/2023 17:48:12
*/


-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS "public"."account";
CREATE TABLE "public"."account" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "open_code" text COLLATE "pg_catalog"."default" NOT NULL,
  "category" int2 NOT NULL,
  "created" int8 NOT NULL,
  "create_rpc" text COLLATE "pg_catalog"."default" NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int8 NOT NULL
)
;
ALTER TABLE "public"."account" OWNER TO "root";
COMMENT ON COLUMN "public"."account"."id" IS '账号ID';
COMMENT ON COLUMN "public"."account"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."account"."open_code" IS '登录账号,如手机号等';
COMMENT ON COLUMN "public"."account"."category" IS '账号类别';
COMMENT ON COLUMN "public"."account"."created" IS '创建时间';
COMMENT ON COLUMN "public"."account"."create_rpc" IS '创建服务';
COMMENT ON COLUMN "public"."account"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."account"."editor" IS '修改人';
COMMENT ON COLUMN "public"."account"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."account" IS '账号';

-- ----------------------------
-- Records of account
-- ----------------------------
BEGIN;
INSERT INTO "public"."account" ("id", "user_id", "open_code", "category", "created", "create_rpc", "edited", "editor", "deleted") VALUES (1, 1, '密码登录', 1, 1691137314, 'public.rpc', 1691137314, 0, 0);
INSERT INTO "public"."account" ("id", "user_id", "open_code", "category", "created", "create_rpc", "edited", "editor", "deleted") VALUES (2, 1, '密码登录', 1, 1691137324, 'public.rpc', 1691137324, 0, 0);
INSERT INTO "public"."account" ("id", "user_id", "open_code", "category", "created", "create_rpc", "edited", "editor", "deleted") VALUES (3, 1, '密码登录', 1, 1691338083, 'public.rpc', 1691338083, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS "public"."group";
CREATE TABLE "public"."group" (
  "id" int8 NOT NULL,
  "parent_id" int8 NOT NULL,
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "code" text COLLATE "pg_catalog"."default" NOT NULL,
  "intro" text COLLATE "pg_catalog"."default" NOT NULL,
  "created" int8 NOT NULL,
  "creator" int8 NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int8 NOT NULL
)
;
ALTER TABLE "public"."group" OWNER TO "root";
COMMENT ON COLUMN "public"."group"."id" IS 'ID';
COMMENT ON COLUMN "public"."group"."parent_id" IS '所属父级用户组ID';
COMMENT ON COLUMN "public"."group"."name" IS '用户组名称';
COMMENT ON COLUMN "public"."group"."code" IS '用户组CODE唯一代码';
COMMENT ON COLUMN "public"."group"."intro" IS '用户组介绍';
COMMENT ON COLUMN "public"."group"."created" IS '创建时间';
COMMENT ON COLUMN "public"."group"."creator" IS '创建人';
COMMENT ON COLUMN "public"."group"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."group"."editor" IS '修改人';
COMMENT ON COLUMN "public"."group"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."group" IS '用户组';

-- ----------------------------
-- Records of group
-- ----------------------------
BEGIN;
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (2, 0, '公司1', 'compare1', '公司1', 1691409406, 1, 0, 1, 0);
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (3, 2, '公司1部门1', 'compare1-1', '公司1部门1', 1691409406, 1, 0, 1, 0);
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (4, 2, '公司1部门2', 'compare1-2', '公司1部门2', 1691409406, 1, 0, 1, 0);
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (5, 4, '公司1部门2-2', 'compare1-2-2', '公司1部门2', 1691409406, 1, 0, 1, 0);
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (1, 0, '超级管理员组', 'adminmmmm    mmmmm', '开发者', 1691409406, 1, 0, 1, 0);
INSERT INTO "public"."group" ("id", "parent_id", "name", "code", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (6, 4, '三级子部门', 'compare1-2-2-1', '公司1部门2', 1691409406, 1, 0, 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS "public"."role";
CREATE TABLE "public"."role" (
  "id" int8 NOT NULL,
  "parent_id" int8 NOT NULL,
  "code" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "intro" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created" int8 NOT NULL,
  "creator" int8 NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int2 NOT NULL
)
;
ALTER TABLE "public"."role" OWNER TO "root";
COMMENT ON COLUMN "public"."role"."id" IS '角色ID';
COMMENT ON COLUMN "public"."role"."parent_id" IS '所属父级角色ID';
COMMENT ON COLUMN "public"."role"."code" IS '角色唯一CODE代码';
COMMENT ON COLUMN "public"."role"."name" IS '角色名称';
COMMENT ON COLUMN "public"."role"."intro" IS '角色介绍';
COMMENT ON COLUMN "public"."role"."created" IS '创建时间';
COMMENT ON COLUMN "public"."role"."creator" IS '创建人';
COMMENT ON COLUMN "public"."role"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."role"."editor" IS '修改人';
COMMENT ON COLUMN "public"."role"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."role" IS '角色';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO "public"."role" ("id", "parent_id", "code", "name", "intro", "created", "creator", "edited", "editor", "deleted") VALUES (1, 1, '超级管理员', '开发者', '开发者', 0, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int8 NOT NULL,
  "state" int2 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "head_img_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "mobile" varchar(11) COLLATE "pg_catalog"."default" NOT NULL,
  "salt" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created" int8 NOT NULL,
  "creator" int8 NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int8 NOT NULL
)
;
ALTER TABLE "public"."user" OWNER TO "root";
COMMENT ON COLUMN "public"."user"."id" IS '用户ID';
COMMENT ON COLUMN "public"."user"."state" IS '用户状态:0=正常,1=禁用';
COMMENT ON COLUMN "public"."user"."name" IS '姓名';
COMMENT ON COLUMN "public"."user"."head_img_url" IS '头像图片地址';
COMMENT ON COLUMN "public"."user"."mobile" IS '手机号码';
COMMENT ON COLUMN "public"."user"."salt" IS '密码加盐';
COMMENT ON COLUMN "public"."user"."password" IS '登录密码';
COMMENT ON COLUMN "public"."user"."created" IS '创建时间';
COMMENT ON COLUMN "public"."user"."creator" IS '创建人';
COMMENT ON COLUMN "public"."user"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."user"."editor" IS '修改人';
COMMENT ON COLUMN "public"."user"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."user" IS '用户';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO "public"."user" ("id", "state", "name", "head_img_url", "mobile", "salt", "password", "created", "creator", "edited", "editor", "deleted") VALUES (1, 1, 'admin', 'https://file.xu756.top/avatar.png', '17337687416', '0', '123456', 0, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group";
CREATE TABLE "public"."user_group" (
  "id" int8 NOT NULL,
  "user_group_id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "created" int8 NOT NULL,
  "creator" int8 NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int8 NOT NULL DEFAULT 0
)
;
ALTER TABLE "public"."user_group" OWNER TO "root";
COMMENT ON COLUMN "public"."user_group"."id" IS 'ID说';
COMMENT ON COLUMN "public"."user_group"."user_group_id" IS '用户组ID';
COMMENT ON COLUMN "public"."user_group"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."user_group"."created" IS '创建时间';
COMMENT ON COLUMN "public"."user_group"."creator" IS '创建人';
COMMENT ON COLUMN "public"."user_group"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."user_group"."editor" IS '修改人';
COMMENT ON COLUMN "public"."user_group"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."user_group" IS '用户组成员';

-- ----------------------------
-- Records of user_group
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_group" ("id", "user_group_id", "user_id", "created", "creator", "edited", "editor", "deleted") VALUES (1, 1, 1, 0, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_role";
CREATE TABLE "public"."user_role" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "role_id" int8 NOT NULL,
  "created" int8 NOT NULL,
  "creator" int8 NOT NULL,
  "edited" int8 NOT NULL,
  "editor" int8 NOT NULL,
  "deleted" int8 NOT NULL
)
;
ALTER TABLE "public"."user_role" OWNER TO "root";
COMMENT ON COLUMN "public"."user_role"."id" IS 'ID';
COMMENT ON COLUMN "public"."user_role"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."user_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."user_role"."created" IS '创建时间';
COMMENT ON COLUMN "public"."user_role"."creator" IS '创建人';
COMMENT ON COLUMN "public"."user_role"."edited" IS '修改时间';
COMMENT ON COLUMN "public"."user_role"."editor" IS '修改人';
COMMENT ON COLUMN "public"."user_role"."deleted" IS '逻辑删除:0=未删除,1=已删除';
COMMENT ON TABLE "public"."user_role" IS '用户角色';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_role" ("id", "user_id", "role_id", "created", "creator", "edited", "editor", "deleted") VALUES (1, 1, 1, 1, 1, 1, 1, 10);
COMMIT;

-- ----------------------------
-- Function structure for getuserinfobyid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."getuserinfobyid"("userid" int8);
CREATE OR REPLACE FUNCTION "public"."getuserinfobyid"("userid" int8)
  RETURNS "public"."user" AS $BODY$
DECLARE
  userInfo "public"."user";
BEGIN
  -- 在 user 表中查询指定 ID 的用户信息，并将结果赋值给 userInfo 变量
  SELECT * INTO userInfo FROM "user" WHERE "id" = userid LIMIT 1;
  
  -- 返回查询结果
  RETURN userInfo;
END
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."getuserinfobyid"("userid" int8) OWNER TO "root";

-- ----------------------------
-- Primary Key structure for table account
-- ----------------------------
ALTER TABLE "public"."account" ADD CONSTRAINT "account_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table group
-- ----------------------------
ALTER TABLE "public"."group" ADD CONSTRAINT "group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table role
-- ----------------------------
ALTER TABLE "public"."role" ADD CONSTRAINT "role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_group
-- ----------------------------
ALTER TABLE "public"."user_group" ADD CONSTRAINT "user_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_role
-- ----------------------------
ALTER TABLE "public"."user_role" ADD CONSTRAINT "user_role_pkey" PRIMARY KEY ("id");
