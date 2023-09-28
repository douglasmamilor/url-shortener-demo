/*
 Navicat Premium Data Transfer

 Source Server         : carbon6-urlshortener
 Source Server Type    : PostgreSQL
 Source Server Version : 130012
 Source Host           : 127.0.0.1:5432
 Source Catalog        : urlshortener
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130012
 File Encoding         : 65001

 Date: 28/09/2023 15:46:53
*/


-- ----------------------------
-- Table structure for url
-- ----------------------------
DROP TABLE IF EXISTS "public"."url";
CREATE TABLE "public"."url" (
  "id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "original_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "short_code" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" time(6) NOT NULL
)
;
ALTER TABLE "public"."url" OWNER TO "urlshortener";

-- ----------------------------
-- Uniques structure for table url
-- ----------------------------
ALTER TABLE "public"."url" ADD CONSTRAINT "short_code" UNIQUE ("short_code");

-- ----------------------------
-- Primary Key structure for table url
-- ----------------------------
ALTER TABLE "public"."url" ADD CONSTRAINT "url_pkey" PRIMARY KEY ("id");
