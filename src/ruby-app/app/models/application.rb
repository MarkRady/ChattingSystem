class Application < ApplicationRecord
    before_create :set_token, :set_chat_count
    validates :Name, presence: true


    private

    def set_token
        self.Token = loop do
            random_token = SecureRandom.urlsafe_base64(nil, false)
            break random_token unless Application.exists?(Token: random_token)
        end
    end

    def set_chat_count
        self.chat_count = 0
    end
end
