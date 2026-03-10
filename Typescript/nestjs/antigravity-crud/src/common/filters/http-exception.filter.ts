import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { Request, Response } from 'express';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = exception.getStatus();
    const exceptionResponse = exception.getResponse();

    const isProduction = process.env.NODE_ENV === 'production';

    const errorBody: Record<string, unknown> = {
      statusCode: status,
      timestamp: new Date().toISOString(),
      path: request.url,
    };

    if (typeof exceptionResponse === 'string') {
      errorBody.message = exceptionResponse;
    } else if (typeof exceptionResponse === 'object') {
      const { message, error } = exceptionResponse as Record<string, unknown>;
      errorBody.message = message;
      errorBody.error = error;
    }

    if (!isProduction && exception.stack) {
      errorBody.stack = exception.stack;
    }

    response.status(status).json(errorBody);
  }
}
