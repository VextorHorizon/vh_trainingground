import { NestFactory } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { AppModule } from './app.module';
import { HttpExceptionFilter } from './common/filters/http-exception.filter';
import { MongooseExceptionFilter } from './common/filters/mongoose-exception.filter';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // Global validation pipe
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      forbidNonWhitelisted: true,
      transform: true,
    }),
  );

  // Global exception filters (order matters: Mongoose filter first, then HTTP)
  app.useGlobalFilters(
    new MongooseExceptionFilter(),
    new HttpExceptionFilter(),
  );

  const configService = app.get(ConfigService);
  const port = configService.get<number>('PORT', 3000);

  await app.listen(port);
  console.log(`🚀 Server running on http://localhost:${port}`);
  console.log(`📦 Environment: ${configService.get<string>('NODE_ENV', 'development')}`);
}
bootstrap();
