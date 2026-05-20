CREATE TABLE IF NOT EXISTS callbacks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    webhook_id UUID NOT NULL REFERENCES webhook(id) ON DELETE CASCADE,
    payload TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    response_code INT NOT NULL,
    attempt_count INT NOT NULL,
    next_retry_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);