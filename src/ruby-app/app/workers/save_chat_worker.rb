class SaveChatWorker
  include Sidekiq::Worker

  def perform
    puts "Hello my world"
    applications = Application.all
    applications.each do |app|
        chatsInCach = Rails.cache.fetch('chatCash_'+app.Token)
        if chatsInCach != nil
            chatsToObject = JSON.parse(chatsInCach)
            chatsToObject.each do |chat|
                Chat.create({Number: chat['Number'], ApplicationId: app.Id, messages_count: 0})
            end
            app.update(chat_count: app.chat_count + chatsToObject.count())
        else
          puts "chatsInCach is empty"
        end
        Rails.cache.delete('chatCash_'+app.Token)
    end
    puts "Queue Done Successfully 💃"

  end
end
