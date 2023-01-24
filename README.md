# Spatial Query Study

| 해당 프로젝트에서는 Spatial Index를 사용해서 공부한 내용을 담았습니다.

![go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![fiber](https://img.shields.io/badge/fiber-00ADD8?logo=go&logoColor=white)
![typescript](https://img.shields.io/badge/typescript-3178C6?logo=typescript&logoColor=white)
![solidjs](https://img.shields.io/badge/solid-2C4F7C?logo=solidjs&logoColor=white)
![mysql](https://img.shields.io/badge/mysql-4479A1?logo=mysql&logoColor=white)

## Spatial Index란?


Spatial Query를 빠르게 수행하기 위한 자료구조입니다.

자세한 내용은 제 블로그([링크](https://velog.io/@noh0907/MySQL-%EC%A2%8C%ED%91%9C-%EB%8D%B0%EC%9D%B4%ED%84%B0-%EA%B0%80%EC%A0%B8%EC%98%A4%EA%B8%B0-Spatial-Index-%ED%99%9C%EC%9A%A9%ED%95%98%EA%B8%B0))에 정리했습니다.

## 스터디 목적

- 데이터 많아지면 성능 저하가 얼마나 생기는지 궁금해서
- 실제 서비스에 적용한다면 체감상 느리게 느껴지는지 궁금해서
- spatial index를 사용한 쿼리와 그렇지 않은 쿼리의 속도 비교

## 결과


### Q. 데이터가 많아지면 성능 저하가 있나요?


**Yes**

당연한 결과입니다.

데이터가 많아지면 데이터베이스에서 처리할 데이터가 많아지기 때문에 속도가 느려집니다

### Q. spatial index를 사용하면 속도 차이가 있나요?


**Yes**

약 11000개의 데이터에서 쿼리를 각각 20번씩 실행했습니다.

거리 차이를 사용한 쿼리의 경우 평균 14.626414ms가 소요됐지만,

spatial query를 사용한 경우 평균 5.526195ms가 소요됐습니다.

### Q. 그럼 spatial index를 사용하면 사용자 입장에서 속도 차이가 체감되나요?


**No** 

프론트엔드에서 지도에 좌표를 보여주는 예제를 통해 테스트한 결과 속도 차이는 느끼지 못했습니다.

2 ~ 3배 차이가 나도, 10ms 단위는 보통 사람이 인지하기 어려운 단위기 때문입니다.

그리고 네트워크 통신이나, asset 다운로드 시간 등 DB외의 요소가 더해지면

실제 쿼리는 1ms 가 나오더라도 사용자에게 보여지는 시간은 10ms 가 넘을 수 있기 때문입니다.

![image](https://user-images.githubusercontent.com/61671343/214256840-0d46f815-90b0-4f33-9f46-fdbbedfc7b20.gif)

**하지만 조건에 따라 Yes가 될 수도 있습니다.**


DB 성능이 안좋은 경우에 spatial query를 사용한 쿼리와 그렇지 않은 쿼리의 속도가 10배정도 차이나는 걸 확인했는데,

spatial query를 사용한 경우 10-20ms 그렇지 않은 경우 190-200ms로 딜레이를 느꼈습니다.

## 예제 실행방법 (docker compose)


위 예시는 아래 커맨드로 로컬에서 바로 테스트할 수 있습니다

### 코드 다운로드


```other
git clone https://github.com/kwanok/spatial-query-study.git
```


### 서버 실행


```other
cd	./spatial-query-study/docker-compose
chmod +x start.bash && ./start.bash
```


### 테스트 사이트 접속 방법


[http://127.0.0.1:5002](http://127.0.0.1:5002) 또는 [localhost:5002](http://localhost:5002)
