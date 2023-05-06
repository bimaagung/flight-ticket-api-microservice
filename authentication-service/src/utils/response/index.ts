export const successResponse = (resData?: any): any => {
  interface Response {
    status: string;
    message?: string;
    data?: any;
  }
  const response: Response = {
    status: 'ok',
    message: 'success',
  };

  if (resData) {
    response.data = resData;
  }

  return response;
};
