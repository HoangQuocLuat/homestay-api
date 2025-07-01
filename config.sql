CREATE TABLE homestays (
  id BIGSERIAL PRIMARY KEY,
  service_id BIGINT,
  host_id BIGINT,
  name VARCHAR(255),
  description TEXT,
  location VARCHAR(255),
  address VARCHAR(255),
  cover_image_url TEXT,
  gallery_images TEXT,
  status INT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);
