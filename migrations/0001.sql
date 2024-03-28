CREATE TABLE IF NOT EXISTS public.helio_credit_nubank
(
  id serial NOT NULL,
  description VARCHAR(255),
  amount integer,
  PRIMARY KEY (id)
);

-- INSERT INTO public.helio_credit_nubank (description, amount) VALUES ('Variedades', 2500);
-- SELECT SUM(amount) from public.helio_credit_nubank;
