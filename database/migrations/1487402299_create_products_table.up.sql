CREATE TABLE products (
  id                serial primary key,
  model             text,
  details           text,
  color             text,
  gender            text,
  features          text[],
  types             text[],
  subtypes          text[],
  sizes             text[],
  primary_image     text,
  images            text[],
  price             numeric,
  discount          numeric,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON products
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

