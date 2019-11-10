class SaveChatWorker
  include Sidekiq::Worker

  def perform
    puts "Hello my world"
    applications = Application.all
    applications.each do |app|
        # puts "Start app:" + app.Token
        # puts 'chatCash_'+app.Token
        chatsInCach = Rails.cache.fetch('chatCash_'+app.Token)
        # puts chatsInCach
        if chatsInCach != nil
            chatsToObject = JSON.parse(chatsInCach)
            # puts "Get :" + chatsToObject

            chatsToObject.each do |chat|
                # puts "Chat is:"+ chat.Number
                Chat.create({Number: chat['Number'], ApplicationId: app.Id, messages_count: 0})
            end
            app.update(chat_count: app.chat_count + chatsToObject.count())
        else
          puts "chatsInCach is empty"
        end
        Rails.cache.delete('chatCash_'+app.Token)
    end
  end
end
