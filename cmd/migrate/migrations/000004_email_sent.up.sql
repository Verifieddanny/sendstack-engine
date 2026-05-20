CREATE TABLE IF NOT EXISTS email_sent (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    receiver_email citext NOT NULL,
    content TEXT NOT NULL,
    subject VARCHAR(255) NOT NULL,
    sender_email citext NOT NULL,
    status VARCHAR(50) NOT NULL,
    delivered_at TIMESTAMPTZ,
    failed_at TIMESTAMPTZ,
    resend_email_id UUID,  -- This is will be changed when I start using useplunk --
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)