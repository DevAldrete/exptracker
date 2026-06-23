-- Schemas for Expense Tracker 

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS expenses (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  cents BIGINT NOT NULL CHECK ( cents >= 0 ),
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS roles (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  description TEXT
);

CREATE TABLE IF NOT EXISTS permissions (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS roles_permissions (
  role_id BIGINT NOT NULL,
  permission_id BIGINT NOT NULL,
  CONSTRAINT fk_roles_permissions_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
  CONSTRAINT fk_roles_permissions_permission FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
  CONSTRAINT pk_roles_permissions PRIMARY KEY(role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS users_roles (
  user_id BIGINT NOT NULL,
  role_id BIGINT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT true,
  CONSTRAINT fk_users_roles_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
  CONSTRAINT fk_users_roles_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT pk_users_roles PRIMARY KEY(user_id, role_id)
);
