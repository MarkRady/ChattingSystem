FROM ruby:2.5.6
RUN apt-get update -qq 

RUN gem install rails -v '5.2.3'
RUN gem update bundler

# Define mountable directories.
VOLUME /root/ruby-app
WORKDIR /root/ruby-app

COPY Gemfile /root/ruby-app/Gemfile
COPY Gemfile.lock /root/ruby-app/Gemfile.lock

RUN bundle install

EXPOSE 3000

# Start the main process.
CMD ["rails", "server", "-b", "0.0.0.0"]
