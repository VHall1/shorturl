# https://docs.docker.com/guides/nodejs/containerize/

FROM node:22.8-alpine

WORKDIR /usr/src/app

RUN --mount=type=bind,source=package.json,target=package.json \
    --mount=type=bind,source=package-lock.json,target=package-lock.json \
    --mount=type=cache,target=/root/.npm \
    npm ci

COPY . .

CMD ["npm", "run", "dev", "--", "-H", "0.0.0.0"]
